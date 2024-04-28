package parser

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/consts"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/repology"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
	"pacstall.dev/webserver/types/pac/parser/git"
	"pacstall.dev/webserver/types/pac/parser/pacsh"
	"pacstall.dev/webserver/types/pac/parser/parallelism/batch"
	"pacstall.dev/webserver/types/pac/parser/parallelism/channels"
)

const PACKAGE_LIST_FILE_NAME = "./packagelist"

func ParseAll() error {
	if err := git.RefreshPrograms(config.GitClonePath, config.GitURL, config.PacstallPrograms.Branch); err != nil {
		return errorx.Decorate(err, "could not update repository 'pacstall-programs'")
	}

	pkgList, err := readKnownPacscriptNames()
	if err != nil {
		return errorx.Decorate(err, "failed to parse packagelist")
	}

	loadedPacscripts, err := parsePacscriptFiles(pkgList)
	if err != nil {
		return errorx.Decorate(err, "failed to parse pacscripts")
	}

	for _, script := range loadedPacscripts {
		computeRequiredBy(script, loadedPacscripts)
	}

	array.SortBy(loadedPacscripts, func(s1, s2 *pac.Script) bool {
		return s1.PackageName < s2.PackageName
	})

	if err := setLastUpdatedAt(loadedPacscripts); err != nil {
		return errorx.Decorate(err, "failed to set last updated at")
	}

	pacstore.Update(loadedPacscripts)
	log.Info("successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(len(pkgList))), len(loadedPacscripts), len(pkgList))

	return nil
}

func readKnownPacscriptNames() ([]string, error) {
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

func parsePacscriptFiles(names []string) ([]*pac.Script, error) {
	if err := pacsh.CreateTempDirectory(config.TempDir); err != nil {
		return nil, errorx.Decorate(err, "failed to create temporary directory")
	}

	log.Info("parsing pacscripts...")
	outChan := batch.Run(int(config.MaxOpenFiles), names, func(pacName string) (*pac.Script, error) {
		out, err := ParsePacscriptFile(config.GitClonePath, pacName)

		if config.Repology.Enabled {
			if err := repology.Sync(&out); err != nil {
				log.Debug("failed to sync %v with repology. Error: %v", pacName, err)
			}
		}

		return &out, err
	})

	return channels.ToSlice(outChan), nil
}

func readPacscriptFile(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	fileName = fmt.Sprintf("%s.%s", name, consts.PACSCRIPT_FILE_EXTENSION)
	scriptPath := path.Join(rootDir, "packages", name, fileName)
	scriptBytes, err = os.ReadFile(scriptPath)

	if err != nil {
		return nil, "", errorx.Decorate(err, "failed to read file '%v'", scriptPath)
	}

	return scriptBytes, fileName, nil
}

func ParsePacscriptFile(programsDirPath, name string) (pac.Script, error) {
	pacshell, filename, err := readPacscriptFile(programsDirPath, name)
	if err != nil {
		return pac.Script{}, errorx.Decorate(err, "failed to read pacscript '%v'", name)
	}

	pacshell = buildCustomFormatScript(pacshell)

	stdout, err := pacsh.ExecBash(config.TempDir, filename, pacshell)
	if err != nil {
		return pac.Script{}, errorx.Decorate(err, "failed to execute pacscript '%v'", name)
	}

	pacscript, err := pacsh.ParsePacOutput(stdout)
	if err != nil {
		return pac.Script{}, errorx.Decorate(err, "failed to parse pacscript '%v'", name)
	}

	return pacscript, nil
}
