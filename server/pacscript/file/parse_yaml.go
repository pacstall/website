package file

import (
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

var ParseYaml = parseYaml

func parseYaml(data []byte, out interface{}) (err error) {
	content := string(data)
	yamlLines := strings.Split(content, "\n")
	content = ""
	for _, line := range yamlLines {
		if len(line) > 0 && line[0] != ' ' && !strings.Contains(line, ":") {
			continue
		}

		content += line + "\n"
	}

	if err = yaml.Unmarshal([]byte(content), &out); err != nil {
		log.Printf("Failed to parse package YAML output\n%v", err)
		log.Fatalln(content)
		return
	}

	return nil
}
