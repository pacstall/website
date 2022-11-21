package parser

import (
	"fmt"
	"os"
	"path"
	"strings"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/git"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/parallelism/batch"
	"pacstall.dev/webserver/parallelism/channels"
	"pacstall.dev/webserver/parser/pacsh"
	"pacstall.dev/webserver/repology"
	"pacstall.dev/webserver/store/pacstore"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

const PACKAGE_LIST_FILE_NAME = "./packagelist"

func ParseAll() {
	if err := git.RefreshPrograms(config.PacstallPrograms.Path, config.PacstallPrograms.URL); err != nil {
		log.Error.Panicln("Could not update repository 'pacstall-programs'", err)
	}

	pkgList, err := readKnownPacscriptNames()
	if err != nil {
		log.Error.Panicln("Failed to parse packagelist", err)
	}

	loadedPacscripts := list.From(parsePacscriptFiles(pkgList)).MapExt(func(p *pac.Script, scripts list.List[*pac.Script]) *pac.Script {
		return computeRequiredBy(*p, scripts)
	}).SortBy(func(s1, s2 *pac.Script) bool {
		return s1.Name < s2.Name
	})

	pacstore.Update(loadedPacscripts)
	log.Info.Printf("Successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(pkgList.Len())), loadedPacscripts.Len(), pkgList.Len())
}

func readKnownPacscriptNames() (list.List[string], error) {
	pkglistPath := path.Join(config.PacstallPrograms.Path, PACKAGE_LIST_FILE_NAME)
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
	if err := pacsh.CreateTempDirectory(config.PacstallPrograms.TempDir); err != nil {
		log.Error.Println(err)
		return nil
	}

	progress := log.NewProgress(len(names), "Parsing pacscripts", "Parsing pacscripts")
	outChan := batch.Run(int(config.PacstallPrograms.MaxOpenFiles), names, func(t string) (*pac.Script, error) {
		out, err := parsePacscriptFile(config.PacstallPrograms.Path, t)
		progress.Describe(fmt.Sprintf("'%v' ok", t))
		progress.Add(1)
		return &out, err
	})

	results := channels.ToSlice(outChan)

	repologySync := repology.NewSyncer(15)
	progressSync := log.NewProgress(len(names), "Syncing with repology", "Syncing with repology")
	for _, result := range results {
		progressSync.Describe(fmt.Sprintf("fetching '%v'", result.Name))
		if err := repologySync(result); err != nil {
			log.Error.Println(err)
		}
		progressSync.Add(1)
	}

	return results
}

func readPacscriptFile(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	fileName = fmt.Sprintf("%v.pacscript", name)
	scriptPath := path.Join(rootDir, "packages", name, fileName)
	scriptBytes, err = os.ReadFile(scriptPath)

	if err != nil {
		log.Error.Printf("Failed to read package pacsh '%v'\n%v", scriptPath, err)
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

	stdout, err := pacsh.ExecBash(config.PacstallPrograms.TempDir, filename, pacshell)
	if err != nil {
		return pac.Script{}, err
	}

	pacscript, err := pacsh.ParsePacOutput(stdout)
	if err != nil {
		return pac.Script{}, fmt.Errorf("failed to parse pacscript %v. err: %v", name, err)
	}

	return pacscript, nil
}
