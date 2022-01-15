package pacpkgs

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"

	"pacstall.dev/website/cfg"
	"pacstall.dev/website/pact"
	"pacstall.dev/website/types"
)

var loadedPackages []types.PackageInfo

func PackageList() []types.PackageInfo {
	return loadedPackages
}

func LoadPackages() {
	if err := pullLatestCommit(); err != nil {
		log.Panicln("Could not update repository 'pacstall-programs'", err)
	}

	pkgList, err := parsePackageList()
	if err != nil {
		log.Panicln("Failed to parse packagelist", err)
	}

	loadedPackages = parsePackages(pkgList)
	log.Printf("Successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPackages))/float64(len(pkgList))), len(loadedPackages), len(pkgList))
}

func ScheduleRefresh(every time.Duration) {
	go func() {
		for {
			time.Sleep(every)
			LoadPackages()
		}
	}()
}

func pullLatestCommit() error {
	cmd := exec.Command("git", "reset", "--hard", "HEAD")
	cmd.Dir = cfg.Config.PacstallPrograms.Path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "fetch")
	cmd.Dir = cfg.Config.PacstallPrograms.Path
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "pull")
	cmd.Dir = cfg.Config.PacstallPrograms.Path
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func parsePackageList() ([]string, error) {
	pkglistPath := path.Join(cfg.Config.PacstallPrograms.Path, "./packagelist")
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

func parsePackages(names []string) []types.PackageInfo {
	parallelism := runtime.NumCPU()
	startedAt := time.Now()
	log.Printf("Running package parsing using %v logical cores\n", parallelism)

	parsedPackages := pact.AsyncBufferedPackageInfo(parallelism, names, parsePackage)

	results := make([]types.PackageInfo, 0)
	for packageInfo := range parsedPackages {
		if packageInfo != nil {
			results = append(results, *packageInfo)
		}
	}

	elapsed := time.Since(startedAt)
	log.Printf("Finished parsing packages after %v. On average, each package took %v", elapsed, elapsed/time.Duration(len(names)))

	return results
}

func parsePackage(name string) *types.PackageInfo {
	log.Printf("Attempting to parse package '%v'\n", name)

	pacscriptName := fmt.Sprintf("%v.pacscript", name)
	scriptPath := path.Join(cfg.Config.PacstallPrograms.Path, "packages", name, pacscriptName)
	scriptBytes, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Printf("Failed to read file '%v'\n%v", scriptPath, err)
		return nil
	}

	tmpPath, err := createTempExecutable(pacscriptName, scriptBytes)
	if err != nil {
		return nil
	}
	defer os.Remove(tmpPath)

	output, err := exec.Command("bash", tmpPath).Output()
	if err != nil {
		log.Printf("Failed to execute '%v'. %v", tmpPath, err)
		return nil
	}

	pkgInfo := rawPackageInfo{}

	if err := json.Unmarshal(output, &pkgInfo); err != nil {
		log.Printf("Failed to parse package JSON output from file '%v'\n%v", tmpPath, err)
		return nil
	}

	return pkgInfo.toPackageInfo()
}

func createTempExecutable(pacscriptName string, content []byte) (string, error) {

	if _, err := os.Stat(cfg.Config.PacstallPrograms.TempDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.Config.PacstallPrograms.TempDir, fs.FileMode(int(0777))); err != nil {
			log.Printf("Failed to create temp dir '%v'\n%v", cfg.Config.PacstallPrograms.TempDir, err)
			return "", err
		}

		log.Printf("Created temp dir '%v'\n", cfg.Config.PacstallPrograms.TempDir)
	}

	tmpFile, err := os.Create(path.Join(cfg.Config.PacstallPrograms.TempDir, pacscriptName))

	if err != nil {
		log.Printf("Failed to create temporary file '%v' in dir '%v'\n", pacscriptName, cfg.Config.PacstallPrograms.TempDir)
		return "", err
	}
	defer tmpFile.Close()
	tmpPath := tmpFile.Name()

	defer func() {
		cmd := exec.Command("chmod", "+rwx", pacscriptName)
		cmd.Dir = cfg.Config.PacstallPrograms.TempDir
		if err := cmd.Run(); err != nil {
			log.Printf("Failed to chmod temporary file '%v' in dir '%v'\n", pacscriptName, cfg.Config.PacstallPrograms.TempDir)
		}
	}()

	if _, err = tmpFile.Write([]byte(BuildJsonScript(string(content)))); err != nil {
		log.Printf("Failed to write to file '%v'\n%v", tmpPath, err)
		return "", err
	}

	if err := tmpFile.Chmod(fs.FileMode(cfg.Config.PacstallPrograms.TempPermissions)); err != nil {
		log.Printf("Failed to chmod file '%v'\n%v", tmpPath, err)
		return "", err
	}

	return tmpPath, nil
}
