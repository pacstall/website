package pacscript

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/hashicorp/go-version"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/pacscript/file"
	"pacstall.dev/webserver/parallelism/batch"
	"pacstall.dev/webserver/parallelism/channels"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

func GetAll() PacscriptList {
	return PacscriptList{
		loadedPacscripts,
	}
}

func LastModified() time.Time {
	return lastModified
}

func Load() {
	if err := pullLatestCommit(); err != nil {
		log.Panicln("Could not update repository 'pacstall-programs'", err)
	}

	pkgList, err := readKnownPacscriptNames()
	if err != nil {
		log.Panicln("Failed to parse packagelist", err)
	}

	loadedPacscripts = list.From(parsePacscriptFiles(pkgList)).MapExt(func(p *pac.Script, scripts list.List[*pac.Script]) *pac.Script {
		return computeRequiredBy(*p, scripts)
	})
	lastModified = time.Now()
	log.Printf("Successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(pkgList.Len())), loadedPacscripts.Len(), pkgList.Len())
}

func ScheduleRefresh(every time.Duration) {
	go func() {
		for {
			time.Sleep(every)
			Load()
		}
	}()
}

func pullLatestCommit() error {
	cmd := exec.Command("git", "reset", "--hard", "HEAD")
	cmd.Dir = config.Config.PacstallPrograms.Path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "fetch")
	cmd.Dir = config.Config.PacstallPrograms.Path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "pull")
	cmd.Dir = config.Config.PacstallPrograms.Path
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func readKnownPacscriptNames() (list.List[string], error) {
	pkglistPath := path.Join(config.Config.PacstallPrograms.Path, "./packagelist")
	bytes, err := os.ReadFile(pkglistPath)
	if err != nil {
		return nil, err
	}

	names := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	for idx := range names {
		names[idx] = strings.TrimSpace(names[idx])
	}

	return names, nil
}

func retryRepologySync(maxRetries int) (*time.Duration, func(*pac.Script) error) {
	baseTime := time.Second * 5
	multiplier := 0.2
	retries := maxRetries
	computedDelay := 1 * time.Second

	return &computedDelay, func(script *pac.Script) error {
		defer func() {
			retries = maxRetries
		}()

		for retries > 0 {

			if multiplier <= 0 {
				multiplier = 1
			}

			computedDelay = baseTime * time.Duration(multiplier)

			retries -= 1
			time.Sleep(computedDelay)

			if retries < maxRetries-1 {
				log.Println("Trying to sync with repology", computedDelay, multiplier)
			}

			if err := fetchAndApplyRepologyInformation(script); err != nil {
				log.Println("Failed to fetch repology information", err)
				multiplier *= 1.5
				continue
			}

			multiplier *= 0.9
			return nil
		}

		return fmt.Errorf("Failed to fetch repology information after %v retries", maxRetries)
	}
}

func benchmark(name string, f func()) {
	start := time.Now()
	f()
	log.Printf("%v took %v", name, time.Since(start))
}

func parsePacscriptFiles(names []string) []*pac.Script {
	if err := file.CreateTempDirectory(config.Config.PacstallPrograms.TempDir); err != nil {
		log.Fatalln(err)
	}

	parseProgress := pb.StartNew(len(names))
	outChan := batch.Run(config.Config.PacstallPrograms.MaxOpenFiles, names, func(t string) (*pac.Script, error) {
		out, err := parsePacscriptFile(config.Config.PacstallPrograms.Path, t)
		parseProgress.Increment()
		return &out, err
	})

	results := channels.ToSlice(outChan)
	parseProgress.Finish()

	_, retrier := retryRepologySync(10)
	log.Println("Syncing with repology...")
	syncProgress := pb.StartNew(len(results))
	for _, result := range results {
		if err := retrier(result); err != nil {
			log.Println(err)
		}
		syncProgress.Increment()
	}
	syncProgress.Finish()

	return results
}

var readFile = os.ReadFile

func fetchAndApplyRepologyInformation(script *pac.Script) (err error) {
	if len(script.Repology) == 0 {
		return
	}

	project, err := fetchRepologyProject(script.Repology)
	if err != nil {
		return
	}

	script.PrettyName = project.PrettyName
	script.LatestVersion = project.Version

	if script.LatestVersion == script.Version || script.Version == "master" || script.Version == "HEAD" || script.Version == "latest" {
		script.UpdateStatus = pac.UpdateStatus.Latest
		return
	}

	current, err := version.NewVersion(script.Version)
	if err != nil {
		err = nil
		script.UpdateStatus = pac.UpdateStatus.Unknown
		return
	}

	latest, err := version.NewVersion(script.LatestVersion)
	if err != nil {
		err = nil
		script.UpdateStatus = pac.UpdateStatus.Unknown
		return
	}

	currentVersionParts := current.Segments64()
	latestVersionParts := latest.Segments64()

	script.UpdateStatus = pac.UpdateStatus.Minor
	if currentVersionParts[0] < latestVersionParts[0] {
		script.UpdateStatus = pac.UpdateStatus.Major
		return
	}

	if len(currentVersionParts) < 2 && len(latestVersionParts) < 2 {
		return
	}

	script.UpdateStatus = pac.UpdateStatus.Patch
	if currentVersionParts[1] < latestVersionParts[1] {
		script.UpdateStatus = pac.UpdateStatus.Minor
		return
	}

	if len(currentVersionParts) < 3 && len(latestVersionParts) < 3 {
		return
	}

	script.UpdateStatus = pac.UpdateStatus.Latest
	if currentVersionParts[2] < latestVersionParts[2] {
		script.UpdateStatus = pac.UpdateStatus.Patch
		return
	}
	return
}

const (
	repologProjectUrl = "https://repology.org/api/v1/project/%s"
)

type repologyProject struct {
	PrettyName string
	Version    string
}

func fetchRepologyProject(search []string) (rpProj repologyProject, err error) {
	getProperty := func(props interface{}, key string) string {
		if props == nil {
			panic("props is nil")
		}

		if v, ok := props.(map[string]interface{})[key]; ok {
			return v.(string)
		}

		panic(fmt.Sprintf("Could not find key %v in %v", key, props))
	}

	project := strings.TrimSpace(strings.Split(search[0], ":")[1])

	resp, err := http.Get(fmt.Sprintf(repologProjectUrl, project))
	if err != nil || resp.StatusCode != 200 {
		return rpProj, fmt.Errorf("(%v) Failed with status %v to fetch repology project via link (%v): %v", project, resp.StatusCode, fmt.Sprintf(repologProjectUrl, project), err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rpProj, fmt.Errorf("Failed to read repology response: %v", err)
	}

	result := make([]interface{}, 0)
	err = json.Unmarshal(body, &result)
	if err != nil {
		os.WriteFile("/tmp/repology.json", body, 0644)
		log.Fatal()
		return rpProj, fmt.Errorf("Failed to unmarshal repology response: %v. \n\n%v\n\n", string(body), err)
	}

	if len(result) == 0 {
		return rpProj, fmt.Errorf("No results for '%v'", project)
	}

	propertyPairs := list.Map(list.From(search[1:]), func(_ int, t string) []string {
		return list.From(strings.Split(t, ":")).Map(func(s string) string {
			return strings.TrimSpace(s)
		})
	})

	foundPackagesRaw := list.From(result).Filter(func(pkg interface{}) bool {
		return list.From(propertyPairs).All(func(pair []string) bool {
			pkgDict := pkg.(map[string]interface{})
			return pkgDict[pair[0]] == pair[1]
		})
	}).SortBy(func(i1, i2 interface{}) bool {
		v1 := i1.(map[string]interface{})["version"]
		v2 := i2.(map[string]interface{})["version"]
		return v1.(string) > v2.(string)
	})

	if foundPackagesRaw.Len() == 0 {
		return rpProj, fmt.Errorf("No results for '%v' after applying search constraints", project)
	}

	rpProj.Version = getProperty(foundPackagesRaw[0], "version")
	rpProj.PrettyName = getProperty(foundPackagesRaw[0], "visiblename")

	if strings.ToLower(rpProj.PrettyName) != rpProj.PrettyName {
		return
	}

	kindaPrettyList := list.From(result).Filter(func(p interface{}) bool {
		visibleName := getProperty(p, "visiblename")
		return strings.ToLower(visibleName) != visibleName
	})

	if kindaPrettyList.IsEmpty() {
		return
	}

	rpProj.PrettyName = getProperty(kindaPrettyList[0], "visiblename")

	veryPrettyList := list.From(kindaPrettyList).Filter(func(p interface{}) bool {
		visibleName := getProperty(p, "visiblename")
		return strings.Contains(visibleName, " ")
	})

	if veryPrettyList.IsEmpty() {
		return
	}

	rpProj.PrettyName = getProperty(veryPrettyList[0], "visiblename")
	return
}

func readPacscriptFile(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	fileName = fmt.Sprintf("%v.pacscript", name)
	scriptPath := path.Join(rootDir, "packages", name, fileName)
	scriptBytes, err = readFile(scriptPath)

	if err != nil {
		log.Printf("Failed to read package file '%v'\n%v", scriptPath, err)
		return
	}

	return scriptBytes, fileName, nil
}

func parsePacscriptFile(programsDirPath, name string) (pac.Script, error) {
	pacsh, filename, err := readPacscriptFile(programsDirPath, name)
	if err != nil {
		return pac.Script{}, err
	}

	pacsh = buildCustomFormatScript(pacsh)

	stdout, err := file.ExecBash(config.Config.PacstallPrograms.TempDir, filename, pacsh)
	if err != nil {
		return pac.Script{}, err
	}

	return file.ParsePacOutput(stdout), nil
}
