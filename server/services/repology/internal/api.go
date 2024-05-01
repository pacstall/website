package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const _USER_AGENT = "Pacstall/WebServer/Exporter"

func GetProjectSearch(projectName string) (RepologyApiProjectSearchResponse, error) {
	var response RepologyApiProjectSearchResponse

	request := http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "https", Host: "repology.org", Path: "/api/v1/projects/" + projectName},
		Header: map[string][]string{
			"Accept":     {"application/json"},
			"User-Agent": {_USER_AGENT},
		},
	}

	resp, err := http.DefaultClient.Do(&request)
	if err != nil {
		return response, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, err
}
