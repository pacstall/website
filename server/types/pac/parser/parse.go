package parser

import (
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
const MAX_GIT_VERSION_CONCURRENCY = 5

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

	log.Info("pacscript parsing done. computing dependency graph")

	for _, script := range loadedPacscripts {
		computeRequiredBy(script, loadedPacscripts)
	}

	array.SortBy(loadedPacscripts, func(s1, s2 *pac.Script) bool {
		return s1.PackageName < s2.PackageName
	})

	log.Info("dependency graph done. setting up git updated-at dates")

	if err := setLastUpdatedAt(loadedPacscripts); err != nil {
		return errorx.Decorate(err, "failed to set last updated at")
	}

	log.Info("updated-at dates done. fetching git versions")

	gitPacscripts := array.Filter(loadedPacscripts, func(it *array.Iterator[*pac.Script]) bool {
		return strings.HasSuffix(it.Value.PackageName, string(types.PACKAGE_TYPE_SUFFIX_GIT))
	})

	channels.Exhaust(batch.Run(MAX_GIT_VERSION_CONCURRENCY, gitPacscripts, func(p *pac.Script) (interface{}, error) {
		err := pacsh.ApplyGitVersion(p)
		return nil, err
	}))

	pacstore.Update(loadedPacscripts)
	log.Info("successfully loaded %v packages from %v pacscripts", len(loadedPacscripts), len(pkgList))

	return nil
}

func readKnownPacscriptNames() ([]string, error) {
	pkglistPath := path.Join(config.GitClonePath, PACKAGE_LIST_FILE_NAME)
	bytes, err := os.ReadFile(pkglistPath)
	if err != nil {
		return nil, err
	}

	names := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	var filteredNames []string

    for idx := range names {
        names[idx] = strings.TrimSpace(names[idx])

		if strings.HasSuffix(names[idx], ":pkgbase") {
            filteredNames = append(filteredNames, strings.TrimSuffix(names[idx], ":pkgbase"))
        } else if !strings.Contains(names[idx], ":") {
            filteredNames = append(filteredNames, names[idx])
        }
    }

    return filteredNames, nil
}

func parsePacscriptFiles(names []string) ([]*pac.Script, error) {
	if err := pacsh.CreateTempDirectory(config.TempDir); err != nil {
		return nil, errorx.Decorate(err, "failed to create temporary directory")
	}

	log.Info("parsing pacscripts...")
	outChan := batch.Run(int(config.MaxOpenFiles), names, func(pacName string) ([]*pac.Script, error) {
		out, err := ParsePacscriptFile(config.GitClonePath, pacName)

		if config.Repology.Enabled {
			for _, script := range out {
		        if err := repology.Sync(script); err != nil {
		            log.Debug("failed to sync %v with repology. Error: %+v", pacName, err)
		        }
		    }
		}

		return out, err
	})
    results := channels.ToSlice(outChan)
    var allScripts []*pac.Script
    for _, scripts := range results {
        allScripts = append(allScripts, scripts...)
    }

    return allScripts, nil
}

func readPacscriptFile(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	scriptPath := path.Join(rootDir, "packages", name, consts.SRCINFO_FILE_EXTENSION)
	scriptBytes, err = os.ReadFile(scriptPath)

	if err != nil {
		return nil, "", errorx.Decorate(err, "failed to read file '%v'", scriptPath)
	}

	return scriptBytes, consts.SRCINFO_FILE_EXTENSION, nil
}

func ParsePacscriptFile(programsDirPath, name string) ([]*pac.Script, error) {
	srcInfoData, _, err := readPacscriptFile(programsDirPath, name)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to read pacscript '%v'", name)
	}

	pacscript, err := pacsh.ParsePacOutput(srcInfoData)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to parse pacscript '%v'", name)
	}

	return pacscript, nil
}
