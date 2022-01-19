package pacpkgs

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
	"pacstall.dev/website/cfg"
	"pacstall.dev/website/types"
)

func PackageList() []types.PackageInfo {
	return loadedPackages
}

func LastModified() time.Time {
	return lastModified
}

func LoadPackages() {
	if err := pullLatestCommit(); err != nil {
		log.Panicln("Could not update repository 'pacstall-programs'", err)
	}

	pkgList, err := parsePackageList()
	if err != nil {
		log.Panicln("Failed to parse packagelist", err)
	}

	loadedPackages = computeRequiredBy(parsePackages(pkgList))
	lastModified = time.Now()
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
	startedAt := time.Now()

	if err := recreateTempDirectory(); err != nil {
		log.Fatalln(err)
	}

	parsedPackages := make(chan *types.PackageInfo)
	guard := make(chan interface{}, cfg.Config.PacstallPrograms.MaxOpenFiles)
	defer close(guard)

	for _, name := range names {
		go func(pkg string) {
			guard <- nil
			parsedPackages <- parsePackage(pkg)
			<-guard
		}(name)
	}

	results := make([]types.PackageInfo, 0)
	for packageInfo := range parsedPackages {
		if packageInfo != nil {
			results = append(results, *packageInfo)
		}

		if len(results) == len(names) {
			close(parsedPackages)
		}
	}

	elapsed := time.Since(startedAt)
	log.Printf("Finished parsing packages after %v. On average, each package took %v", elapsed, elapsed/time.Duration(len(names)))

	return results
}

func parsePackage(name string) *types.PackageInfo {
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
		bytes, _ := os.ReadFile(tmpPath)
		log.Printf("Failed to execute '%v'. %v\n%v", tmpPath, err, string(bytes))
		return nil
	}

	content := string(output)
	yamlLines := strings.Split(content, "\n")
	content = ""
	for _, line := range yamlLines {
		if len(line) > 0 && line[0] != ' ' && !strings.Contains(line, ":") {
			continue
		}

		content += line + "\n"
	}

	rawPkgInfo := rawPackageInfo{}

	if err := yaml.Unmarshal([]byte(content), &rawPkgInfo); err != nil {
		log.Printf("Failed to parse package YAML output from file '%v'\n%v", tmpPath, err)
		log.Fatalln(content)
		return nil
	}

	pkgInfo := rawPkgInfo.toPackageInfo()

	if err := RepairPackageInfo(&pkgInfo); err != nil {
		log.Printf("Failed to repair package info type for '%v'\n", name)
		return nil
	}

	return &pkgInfo
}

func recreateTempDirectory() error {
	if _, err := os.Stat(cfg.Config.PacstallPrograms.TempDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.Config.PacstallPrograms.TempDir, fs.FileMode(int(0777))); err != nil {
			log.Printf("Failed to create temp dir '%v'\n%v", cfg.Config.PacstallPrograms.TempDir, err)
			return err
		}

		log.Printf("Created fresh temp dir '%v'\n", cfg.Config.PacstallPrograms.TempDir)
	} else {
		if err := os.RemoveAll(cfg.Config.PacstallPrograms.TempDir); err != nil {
			log.Printf("Failed to remove existing temp dir '%v'\n", cfg.Config.PacstallPrograms.TempDir)
			return err
		}

		log.Printf("Removed existing temp dir '%v'\n", cfg.Config.PacstallPrograms.TempDir)
		return recreateTempDirectory()
	}

	return nil
}

func createTempExecutable(pacscriptName string, content []byte) (string, error) {
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

	if _, err = tmpFile.Write([]byte(buildYamlScript(string(content)))); err != nil {
		log.Printf("Failed to write to file '%v'\n%v", tmpPath, err)
		return "", err
	}

	if err := tmpFile.Chmod(fs.FileMode(int(0777))); err != nil {
		log.Printf("Failed to chmod file '%v'\n%v", tmpPath, err)
		return "", err
	}

	return tmpPath, nil
}
