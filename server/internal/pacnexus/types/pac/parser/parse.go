package parser

import (
	"os"
	"path"
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/internal/pacnexus/config"
	"pacstall.dev/webserver/internal/pacnexus/consts"
	"pacstall.dev/webserver/internal/pacnexus/types/pac"
	"pacstall.dev/webserver/internal/pacnexus/types/pac/pacstore"
	"pacstall.dev/webserver/internal/pacnexus/types/pac/parser/git"
	"pacstall.dev/webserver/internal/pacnexus/types/pac/parser/pacsh"
	"pacstall.dev/webserver/pkg/common/array"
	"pacstall.dev/webserver/pkg/common/log"
	"pacstall.dev/webserver/pkg/common/pacsight"
	"pacstall.dev/webserver/pkg/common/parallelism/batch"
	"pacstall.dev/webserver/pkg/common/parallelism/channels"
	"pacstall.dev/webserver/pkg/common/types"
)

const PACKAGE_LIST_FILE_NAME = "./packagelist"
const MAX_GIT_VERSION_CONCURRENCY = 5

func ParseAll(pacsightRpc *pacsight.PacsightRpcService) error {
	if err := git.RefreshPrograms(config.PacstallPrograms.GitClonePath, config.PacstallPrograms.GitURL, config.PacstallPrograms.Branch); err != nil {
		return errorx.Decorate(err, "could not update repository 'pacstall-programs'")
	}

	pkgList, err := readKnownPacscriptNames()
	if err != nil {
		return errorx.Decorate(err, "failed to parse packagelist")
	}

	loadedPacscripts, err := parsePacscriptFiles(pkgList, pacsightRpc)
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
	log.Info("successfully loaded %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(len(pkgList))), len(loadedPacscripts), len(pkgList))

	return nil
}

func readKnownPacscriptNames() ([]string, error) {
	pkglistPath := path.Join(config.PacstallPrograms.GitClonePath, PACKAGE_LIST_FILE_NAME)
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

func parsePacscriptFiles(names []string, pacsightRpc *pacsight.PacsightRpcService) ([]*pac.Script, error) {
	if err := pacsh.CreateTempDirectory(config.PacstallPrograms.TempDir); err != nil {
		return nil, errorx.Decorate(err, "failed to create temporary directory")
	}

	log.Info("parsing pacscripts...")
	outChan := batch.Run(int(config.PacstallPrograms.MaxOpenFiles), names, func(pacName string) (*pac.Script, error) {
		out, err := ParsePacscriptFile(config.PacstallPrograms.GitClonePath, pacName)

		if config.Repology.Enabled && len(out.Repology) > 0 {
			project, err := pacsightRpc.GetRepologyProject(out.Repology)
			if err != nil {
				log.Debug("failed to get repology project %v. Error: %+v", out.Repology, err)
			} else {
				if err := UpdateScriptVersion(project, out); err != nil {
					log.Warn("failed to update script version %v with repology. Error: %+v", pacName, err)
				}

				out.LatestVersion = &project.Version

			}
		}

		return out, err
	})

	return channels.ToSlice(outChan), nil
}

func readPacscriptFile(rootDir, name string) (scriptBytes []byte, fileName string, err error) {
	scriptPath := path.Join(rootDir, "packages", name, consts.SRCINFO_FILE_EXTENSION)
	scriptBytes, err = os.ReadFile(scriptPath)

	if err != nil {
		return nil, "", errorx.Decorate(err, "failed to read file '%v'", scriptPath)
	}

	return scriptBytes, consts.SRCINFO_FILE_EXTENSION, nil
}

func ParsePacscriptFile(programsDirPath, name string) (*pac.Script, error) {
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
