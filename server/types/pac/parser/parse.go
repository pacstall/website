package parser

import (
	"fmt"
	"os"
	"path"
	"strings"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/repology"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
	"pacstall.dev/webserver/types/pac/parser/git"
	"pacstall.dev/webserver/types/pac/parser/pacsh"
	"pacstall.dev/webserver/types/pac/parser/parallelism/batch"
	"pacstall.dev/webserver/types/pac/parser/parallelism/channels"
)

const PACKAGE_LIST_FILE_NAME = "./packagelist"

func ParseAll() {
	if err := git.RefreshPrograms(config.GitClonePath, config.GitURL); err != nil {
		log.Fatal("Could not update repository 'pacstall-programs'. %v", err)
	}

	pkgList, err := readKnownPacscriptNames()
	if err != nil {
		log.Fatal("Failed to parse packagelist. %v", err)
	}

	loadedPacscripts := list.From(parsePacscriptFiles(pkgList)).MapExt(func(p *pac.Script, scripts list.List[*pac.Script]) *pac.Script {
		return computeRequiredBy(*p, scripts)
	}).SortBy(func(s1, s2 *pac.Script) bool {
		return s1.Name < s2.Name
	})

	pacstore.Update(loadedPacscripts)
	log.Info("Successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(pkgList.Len())), loadedPacscripts.Len(), pkgList.Len())
	log.Notify("Successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(pkgList.Len())), loadedPacscripts.Len(), pkgList.Len())
}

func readKnownPacscriptNames() (list.List[string], error) {
	pkglistPath := path.Join(config.GitClonePath, PACKAGE_LIST_FILE_NAME)
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

func parsePacscriptFiles(names []string) []*pac.Script {
	if err := pacsh.CreateTempDirectory(config.TempDir); err != nil {
		log.Error("Failed to create temporary directory. %v", err)
		return nil
	}

	log.Info("Parsing pacscripts...")
	outChan := batch.Run(int(config.MaxOpenFiles), names, func(pacName string) (*pac.Script, error) {
		out, err := parsePacscriptFile(config.GitClonePath, pacName)
		if err != nil {
			log.Warn("Failed to parse %v. err: %v", pacName, err)
		}
		return &out, err
	})

	results := channels.ToSlice(outChan)

	repologySync := repology.NewSyncer(15)
	log.Info("Syncing pacscripts with repology...")

	for _, result := range results {
		if err := repologySync(result); err != nil {
			log.Warn("Failed to sync %v. err: %v", result.Name, err)
		}
	}

	return results
}

func readPacscriptFile(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	fileName = fmt.Sprintf("%v.pacscript", name)
	scriptPath := path.Join(rootDir, "packages", name, fileName)
	scriptBytes, err = os.ReadFile(scriptPath)

	if err != nil {
		log.Error("Failed to read package pacsh '%v'\n%v", scriptPath, err)
		return
	}

	return scriptBytes, fileName, nil
}

func parsePacscriptFile(programsDirPath, name string) (pac.Script, error) {
	pacshell, filename, err := readPacscriptFile(programsDirPath, name)
	if err != nil {
		return pac.Script{}, err
	}

	pacshell = buildCustomFormatScript(pacshell)

	stdout, err := pacsh.ExecBash(config.TempDir, filename, pacshell)
	if err != nil {
		return pac.Script{}, err
	}

	pacscript, err := pacsh.ParsePacOutput(stdout)
	if err != nil {
		return pac.Script{}, fmt.Errorf("failed to parse pacscript %v. err: %v", name, err)
	}

	return pacscript, nil
}
