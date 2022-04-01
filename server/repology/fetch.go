package repology

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"pacstall.dev/webserver/types/list"
)

var repologyCache = make(map[string][]repologyRawProject)
var cachedUpdatedAt = time.Now()

func fetchRaw(project string) ([]repologyRawProject, error) {
	if cachedUpdatedAt.After(time.Now().Add(-time.Minute*5)) && repologyCache[project] != nil {
		return repologyCache[project], nil
	}

	resp, err := http.Get(fmt.Sprintf(repologProjectUrl, project))
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("(%v) Failed with status %v to fetch repology project via link (%v): %v", project, resp.StatusCode, fmt.Sprintf(repologProjectUrl, project), err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read repology response: %v", err)
	}

	result := make([]repologyRawProject, 0)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal repology response: %v. \n\n%v\n\n", string(body), err)
	}

	repologyCache[project] = result

	if len(result) == 0 {
		return nil, fmt.Errorf("No results for '%v'", project)
	}

	return result, nil
}

func getProperty(project repologyRawProject, property string) string {
	if project[property] == nil {
		return ""
	}

	return project[property].(string)
}

func getSliceProperty(project repologyRawProject, property string) []string {
	if project[property] == nil {
		return nil
	}

	return list.Map(project[property].([]interface{}), func(i int, t interface{}) string {
		return fmt.Sprintf("%v", t)
	})
}

func fetchRepologyProject(search []string) (rpProj repologyProject, err error) {
	project := strings.TrimSpace(strings.Split(search[0], ":")[1])
	result, err := fetchRaw(project)
	if err != nil {
		return
	}

	propertyPairs := list.Map(list.From(search[1:]), func(_ int, t string) []string {
		return list.From(strings.Split(t, ":")).Map(func(s string) string {
			return strings.TrimSpace(s)
		})
	})

	foundPackagesRaw := list.Map(list.From(result).Filter(func(pkg repologyRawProject) bool {
		return list.From(propertyPairs).All(func(pair []string) bool {
			return pkg[pair[0]] == pair[1] && strings.ContainsAny(pkg["version"].(string), "1234567890")
		})
	}), func(i int, t repologyRawProject) repologySemiRawProject {
		return repologySemiRawProject{
			Name:        getProperty(t, "name"),
			Version:     getProperty(t, "version"),
			VisibleName: getProperty(t, "visiblename"),
			Summary:     getProperty(t, "summary"),
			Repo:        getProperty(t, "repo"),
			Status:      getProperty(t, "status"),
			SrcName:     getProperty(t, "srcname"),
			BinName:     getProperty(t, "binname"),
			SubRepo:     getProperty(t, "subrepo"),
			Licenses:    getSliceProperty(t, "licenses"),
			OrigVersion: getProperty(t, "origversion"),
			Maintainers: getSliceProperty(t, "maintainers"),
			Categories:  getSliceProperty(t, "categories"),
		}
	}).Filter(func(pkg repologySemiRawProject) bool {
		return pkg.Status != "incorrect"
	}).SortBy(func(p1, p2 repologySemiRawProject) bool {
		v1HasNumbers := strings.ContainsAny(p1.Version, "0123456789")
		v2HasNumbers := strings.ContainsAny(p2.Version, "0123456789")

		if v1HasNumbers && !v2HasNumbers {
			return true
		} else if !v1HasNumbers && v2HasNumbers {
			return false
		} else if !v1HasNumbers && !v2HasNumbers {
			return true
		}

		v1, err := version.NewVersion(p1.Version)
		if err != nil {
			return false
		}

		v2, err := version.NewVersion(p2.Version)
		if err != nil {
			return true
		}

		return v1.GreaterThan(v2)
	}).SortBy(func(rsrp1, rsrp2 repologySemiRawProject) bool {
		return !(rsrp1.Status == "newest" && rsrp2.Status != "newest")
	})

	if foundPackagesRaw.Len() == 0 {
		return rpProj, fmt.Errorf("No results for '%v' after applying search constraints", project)
	}

	rpProj.Version = foundPackagesRaw[0].Version
	rpProj.PrettyName = foundPackagesRaw[0].VisibleName

	if strings.ToLower(rpProj.PrettyName) != rpProj.PrettyName {
		return
	}

	kindaPrettyList := foundPackagesRaw.Filter(func(p repologySemiRawProject) bool {
		return strings.ToLower(p.VisibleName) != p.VisibleName
	})

	if kindaPrettyList.IsEmpty() {
		return
	}

	rpProj.PrettyName = kindaPrettyList[0].VisibleName

	veryPrettyList := list.From(kindaPrettyList).Filter(func(p repologySemiRawProject) bool {
		return strings.Contains(p.VisibleName, " ")
	})

	if veryPrettyList.IsEmpty() {
		return
	}

	rpProj.PrettyName = veryPrettyList[0].VisibleName
	return
}
