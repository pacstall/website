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
	"pacstall.dev/webserver/services/parser/pacsh"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/service"
	"pacstall.dev/webserver/utils/git"
	"pacstall.dev/webserver/utils/parallelism/batch"
	"pacstall.dev/webserver/utils/parallelism/channels"
)

const PACKAGE_LIST_FILE_NAME = "./packagelist"

type ParserService struct {
	programsConfiguration     config.PacstallProgramsConfiguration
	serverConfiguration       config.ServerConfiguration
	repologyConfiguration     config.RepologyConfiguration
	repologyService           service.RepologyService
	gitVersionResolver        service.GitVersionResolver
	packageCacheService       service.PackageCacheService
	packageLastUpdatedService service.PackageLastUpdatedService
}

func New(
	programsConfiguration config.PacstallProgramsConfiguration,
	serverConfiguration config.ServerConfiguration,
	repologyConfiguration config.RepologyConfiguration,
	repologyService service.RepologyService,
	gitVersionResolver service.GitVersionResolver,
	packageCacheService service.PackageCacheService,
	packageLastUpdatedService service.PackageLastUpdatedService,

) *ParserService {
	s := &ParserService{}

	s.programsConfiguration = programsConfiguration
	s.serverConfiguration = serverConfiguration
	s.repologyConfiguration = repologyConfiguration
	s.repologyService = repologyService
	s.gitVersionResolver = gitVersionResolver
	s.packageCacheService = packageCacheService
	s.packageLastUpdatedService = packageLastUpdatedService

	return s
}

func (s *ParserService) ParseAll() error {

	if err := git.RefreshPrograms(s.programsConfiguration.ClonePath, s.programsConfiguration.RepositoryUrl, s.programsConfiguration.Branch); err != nil {
		return errorx.Decorate(err, "could not update repository 'pacstall-programs'")
	}

	pkgList, err := s.loadPackageNamesFromFile()
	if err != nil {
		return errorx.Decorate(err, "failed to parse packagelist")
	}

	loadedPacscripts, err := s.parsePacscriptFiles(pkgList)
	if err != nil {
		return errorx.Decorate(err, "failed to parse pacscripts")
	}

	for _, script := range loadedPacscripts {
		computeRequiredBy(script, loadedPacscripts)
	}

	array.SortBy(loadedPacscripts, func(s1, s2 *pac.Script) bool {
		return s1.PackageName < s2.PackageName
	})

	if err := s.setLastUpdatedAt(loadedPacscripts, s.programsConfiguration.ClonePath); err != nil {
		return errorx.Decorate(err, "failed to set last updated at")
	}

	s.packageCacheService.Update(loadedPacscripts)
	log.Info("successfully parsed %v (%v / %v) packages", types.Percent(float64(len(loadedPacscripts))/float64(len(pkgList))), len(loadedPacscripts), len(pkgList))

	return nil
}

func (s *ParserService) loadPackageNamesFromFile() ([]string, error) {
	pkglistPath := path.Join(s.programsConfiguration.ClonePath, PACKAGE_LIST_FILE_NAME)
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

func (s *ParserService) parsePacscriptFiles(names []string) ([]*pac.Script, error) {
	if err := pacsh.CreateTempDirectory(s.serverConfiguration.TempDir); err != nil {
		return nil, errorx.Decorate(err, "failed to create temporary directory")
	}

	log.Info("started batch job to parse pacscripts...")
	outChan := batch.Run(s.serverConfiguration.MaxOpenFiles, names, func(pacName string) (*pac.Script, error) {
		out, err := s.ParsePacscriptFile(s.programsConfiguration.ClonePath, pacName)

		if s.repologyConfiguration.Enabled {
			if err := s.repologyService.Sync(&out); err != nil {
				log.Debug("failed to sync %v with repology. Error: %v", pacName, err)
			}
		}

		return &out, err
	})

	return channels.ToSlice(outChan), nil
}

func readPacscriptFile(rootDir, name string) (script string, fileName string, err error) {
	fileName = fmt.Sprintf("%s.%s", name, consts.PACSCRIPT_FILE_EXTENSION)
	scriptPath := path.Join(rootDir, "packages", name, fileName)
	scriptBytes, err := os.ReadFile(scriptPath)

	if err != nil {
		return "", "", errorx.Decorate(err, "failed to read file '%v'", scriptPath)
	}

	return string(scriptBytes), fileName, nil
}

func (s *ParserService) ParsePacscriptFile(programsDirPath, name string) (pac.Script, error) {
	pacshell, filename, err := readPacscriptFile(programsDirPath, name)
	if err != nil {
		return pac.Script{}, errorx.Decorate(err, "failed to read pacscript '%v'", name)
	}

	pacshell = buildCustomFormatScript(pacshell)

	stdout, err := pacsh.ExecBash(s.serverConfiguration.TempDir, filename, pacshell)
	if err != nil {
		return pac.Script{}, errorx.Decorate(err, "failed to execute pacscript '%v'", name)
	}

	pacscript, err := pacsh.ParsePacOutput(s.gitVersionResolver, stdout)
	if err != nil {
		return pac.Script{}, errorx.Decorate(err, "failed to parse pacscript '%v'", name)
	}

	return pacscript, nil
}
