package parser

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/parser/pacsh"
)

type packageLastUpdatedTuple struct {
	packageName string
	lastUpdated time.Time
}

func getPackageLastUpdatedTuples() ([]packageLastUpdatedTuple, error) {
	wordingDirectoryAbsolute, err := os.Getwd()
	if err != nil {
		return nil, errorx.Decorate(err, "failed to get absolute path to wording directory")
	}

	programsPath := path.Join(wordingDirectoryAbsolute, config.GitClonePath)
	script := fmt.Sprintf(`
	cd %v
	for i in ./packages/*/*.pacscript; do echo $i; git log -1 --pretty=\"%%at\" $i; done
	`, programsPath)

	outputBytes, err := pacsh.ExecBash(programsPath, "last_updated.sh", []byte(script))
	if err != nil {
		return nil, errorx.Decorate(err, "failed to get last updated git output")
	}

	output := string(outputBytes)
	lines := strings.Split(output, "\n")
	lines = lines[:len(lines)-1] // Remove last empty line
	tuples := make([]packageLastUpdatedTuple, 0)

	for i := 0; i < len(lines)-1; i += 2 {
		packagePath := lines[i]
		lastUpdatedString := lines[i+1] // Unix time

		// Remove quotes
		lastUpdatedString = lastUpdatedString[1 : len(lastUpdatedString)-1]

		packageNameWithExtension := path.Base(packagePath)
		packageName := strings.TrimSuffix(packageNameWithExtension, ".pacscript")

		if packageName == "" || strings.HasPrefix(packageName, "-") {
			return nil, errorx.IllegalState.New("failed to parse package name from package path '%v'", packagePath)
		}

		lastUpdatedUnixTime, err := strconv.ParseInt(lastUpdatedString, 10, 64)
		if err != nil {
			return nil, errorx.Decorate(err, "failed to parse '%v' as int64", lastUpdatedString)
		}

		lastUpdated := time.Unix(lastUpdatedUnixTime, 0).UTC()

		if lastUpdated.Year() < 2000 {
			return nil, errorx.IllegalState.New("failed to parse last updated time for package '%v'. Given date is %v", packagePath, lastUpdatedString)
		}

		tuples = append(tuples, packageLastUpdatedTuple{
			packageName: packageName,
			lastUpdated: lastUpdated,
		})
	}

	return tuples, nil
}

func setLastUpdatedAt(packages []*pac.Script) error {
	lastUpdatedTuples, err := getPackageLastUpdatedTuples()
	if err != nil {
		return errorx.Decorate(err, "failed to get package last updated tuples")
	}

	for _, pkg := range packages {
		if tuple, err := array.FindBy(lastUpdatedTuples, func(tuple packageLastUpdatedTuple) bool {
			return tuple.packageName == pkg.PackageName
		}); err == nil {
			pkg.LastUpdatedAt = tuple.lastUpdated
		} else {
			log.Warn("failed to set 'LastUpdatedAt' for package %#v. err: %+v", pkg, err)
		}
	}

	return nil
}
