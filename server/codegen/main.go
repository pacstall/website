package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

type GenFile struct {
	path    string
	replace map[string]string
	imports []string
	content string
}

func findFiles() []string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panicln("Failed to get cwd", err)
	}

	return findFilesInDir(cwd)
}

func findFilesInDir(dir string) []string {
	out := make([]string, 0)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Panicln("Failed to list dir files", dir, err)
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Name())

		if file.IsDir() {
			innerFiles := findFilesInDir(filePath)
			out = append(out, innerFiles...)
		} else if strings.HasSuffix(filePath, ".go") && !strings.HasSuffix(filePath, "_generated.go") {
			out = append(out, filePath)
		}
	}

	return out
}

func validateLine(file *GenFile, lineNo int, filePath string) {
	if file == nil {
		log.Fatalf("Cannot parse file '%v' at line %v. Found codegen instruction before '@generate' was specified.", filePath, lineNo+1)
	}
}

func parseGenFile(filePath string) []*GenFile {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicln("Failed to read file", filePath, err)
	}

	content := string(bytes)
	lines := strings.Split(content, "\n")
	filePathParts := strings.Split(filePath, "/")
	dirPath := strings.Join(filePathParts[:len(filePathParts)-1], "/")
	fileName := filePathParts[len(filePathParts)-1]

	genFiles := make([]*GenFile, 0)
	var genFile *GenFile
	for lineNo, line := range lines {
		if strings.HasPrefix(line, "// @generate") {
			genFileNameSuffix := strings.TrimSpace(strings.Split(line, "@generate")[1])
			genFileName := strings.Replace(fileName, ".go", fmt.Sprintf("_%v_generated.go", genFileNameSuffix), 1)

			genFile = &GenFile{
				path:    path.Join(dirPath, genFileName),
				replace: make(map[string]string),
				imports: make([]string, 0),
				content: content,
			}

			genFiles = append(genFiles, genFile)
		}

		if strings.HasPrefix(line, "// @import") {
			validateLine(genFile, lineNo, filePath)

			genFile.imports = append(genFile.imports, strings.TrimSpace(strings.Split(line, "@import")[1]))
		}

		if strings.HasPrefix(line, "// @replace") {
			validateLine(genFile, lineNo, filePath)

			rhs := strings.TrimSpace(strings.Split(line, "@replace")[1])
			replacements := strings.Split(rhs, ">>")
			if len(replacements) != 2 {
				log.Fatalf("Failed to parse line '%v' from file '%v' line %v. Must follow this structure: '// @replace A >> B'\n", line, filePath, lineNo)
			}

			genFile.replace[strings.TrimSpace(replacements[0])] = strings.TrimSpace(replacements[1])
		}
	}

	return genFiles
}

func generateFile(file *GenFile) {
	if _, err := os.Stat(file.path); err == nil {
		if err = os.Remove(file.path); err != nil {
			log.Fatalln("Failed to remove file", file.path, err)
		}
	}

	importContent := ""
	for _, imprt := range file.imports {
		importContent = fmt.Sprintf("%v\nimport \"%v\"", importContent, imprt)
	}

	for toReplace, replacement := range file.replace {
		file.content = strings.ReplaceAll(file.content, fmt.Sprintf("type %v =", toReplace), fmt.Sprintf("// type %v =", replacement))
		file.content = strings.ReplaceAll(file.content, toReplace, replacement)
	}

	lines := strings.Split(file.content, "\n")
	packageDeclaration := lines[0]
	restOfFile := strings.Join(lines[1:], "\n")

	generated := packageDeclaration
	generated = fmt.Sprintf("%v\n\n// This file is generated. Do not edit!", generated)
	generated = fmt.Sprintf("%v\n\n%v\n%v", generated, importContent, restOfFile)

	if err := os.WriteFile(file.path, []byte(generated), 0755); err != nil {
		log.Panicln("Failed to write file", file.path, err)
	}
}

func main() {
	wg := sync.WaitGroup{}
	for _, file := range findFiles() {
		for _, genFile := range parseGenFile(file) {
			wg.Add(1)

			go func(genFile *GenFile) {
				generateFile(genFile)
				wg.Done()
			}(genFile)
		}
	}

	wg.Wait()
	log.Println("Done.")
}
