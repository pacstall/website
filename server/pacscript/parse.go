package pacscript

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/pacscript/file"
	"pacstall.dev/webserver/parallelism/batch"
	"pacstall.dev/webserver/parallelism/channels"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/list"
)

func GetAll() PackageList {
	return PackageList{
		loadedPackages,
	}
}

func LastModified() time.Time {
	return lastModified
}

func Load() {
	if err := pullLatestCommit(); err != nil {
		log.Panicln("Could not update repository 'pacstall-programs'", err)
	}

	pkgList, err := parsePackageList()
	if err != nil {
		log.Panicln("Failed to parse packagelist", err)
	}

	loadedPackages = computeRequiredBy(parsePackages(pkgList))
	lastModified = time.Now()
	log.Printf("Successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPackages))/float64(pkgList.Len())), loadedPackages.Len(), pkgList.Len())
}

var ScheduleRefresh = func(every time.Duration) {
	go func() {
		for {
			time.Sleep(every)
			Load()
		}
	}()
}

var pullLatestCommit = func() error {
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

var parsePackageList = func() (list.List[string], error) {
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

var parsePackages = func(names []string) []*types.Pacscript {
	startedAt := time.Now()

	if err := file.CreateTempDirectory(config.Config.PacstallPrograms.TempDir); err != nil {
		log.Fatalln(err)
	}

	outChan := batch.Run(config.Config.PacstallPrograms.MaxOpenFiles, names, func(t string) (*types.Pacscript, error) {
		result, err := parsePackage(config.Config.PacstallPrograms.Path, t)
		return &result, err
	})

	results := channels.ToSlice(outChan)

	elapsed := float32(time.Since(startedAt)) / float32(time.Second)
	each := float32(time.Since(startedAt)) / float32(time.Duration(len(names))) / float32(time.Millisecond)
	log.Printf("Finished parsing packages after %.2fs. On average, each package took %.2fms", elapsed, each)

	return results
}

var readFile = os.ReadFile

func readPacscript(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	fileName = fmt.Sprintf("%v.pacscript", name)
	scriptPath := path.Join(rootDir, "packages", name, fileName)
	scriptBytes, err = readFile(scriptPath)

	if err != nil {
		log.Printf("Failed to read package file '%v'\n%v", scriptPath, err)
		return
	}

	return scriptBytes, fileName, nil
}

var parsePackage = func(programsDirPath, name string) (pacscript types.Pacscript, err error) {
	pacsh, filename, err := readPacscript(programsDirPath, name)
	if err != nil {
		return
	}

	pacsh = buildYamlScript(pacsh)
	stdout, err := file.ExecBash(config.Config.PacstallPrograms.TempDir, filename, pacsh)

	rawPkgInfo := rawPacscript{}
	if err = file.ParseYaml(stdout, &rawPkgInfo); err != nil {
		log.Printf("Failed to parse package '%v'\n%v", name, err)
		return
	}

	pacscript = rawPkgInfo.toPacscript()

	if err = RepairPacscript(&pacscript); err != nil {
		log.Printf("Failed to repair package info type for '%v'\n", name)
		return
	}

	return pacscript, nil
}
