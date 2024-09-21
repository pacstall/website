package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/joomcode/errorx"
)

const _USER_AGENT = "Pacstall/WebServer/Exporter"

func getProjectSearch(projectName string) (RepologyApiProjectSearchResponse, error) {
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
		return response, errorx.Decorate(err, "http request failed %+v", request)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, errorx.Decorate(err, "failed to read request body")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return response, errorx.RejectedOperation.New("http request failed with status code %v. body \n%v\n", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, errorx.Decorate(err, "failed to unmarshal response body '%v'", string(body))
	}

	return response, err
}
