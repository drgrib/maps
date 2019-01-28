package maptemplate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type MapFile struct {
	FileName    string
	MapTemplate MapTemplate
}

func (mapFile MapFile) Write() error {
	content, err := mapFile.MapTemplate.Parse()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(mapFile.FileName, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func NewMapFile(keytype, valuetype string) (*MapFile, error) {
	packageName, err := inferPackage()
	if err != nil {
		return nil, err
	}

	mapTemplate := NewMapTemplate(packageName, keytype, valuetype)
	fileName := getDefaultFilename(mapTemplate)

	mapFile := MapFile{
		FileName:    fileName,
		MapTemplate: mapTemplate,
	}

	return &mapFile, nil
}

func getDefaultFilename(mapTemplate MapTemplate) string {
	return fmt.Sprintf("map_%s_%s.go", strings.ToLower(mapTemplate.KeySuffix), strings.ToLower(mapTemplate.ValueSuffix))
}

func inferPackage() (string, error) {
	goFiles, err := getWdGoFiles()
	if err != nil {
		return "", err
	}

	if len(goFiles) == 0 {
		return getCurrentDir()
	}

	firstGoFilePath := goFiles[0]
	packageName, err := getFilePackage(firstGoFilePath)
	if err != nil {
		return "", err
	}

	if packageName == "" {
		return getCurrentDir()
	}

	return packageName, nil
}

func getCurrentDir() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := filepath.Base(p)

	return dir, nil
}

func getWdGoFiles() ([]string, error) {
	goFiles, err := filepath.Glob("*.go")
	if err != nil {
		return nil, err
	}

	return goFiles, nil
}

func getFilePackage(p string) (string, error) {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}

	r, err := regexp.Compile(`package (\w+)`)
	if err != nil {
		return "", err
	}

	matches := r.FindStringSubmatch(string(b))
	if len(matches) <= 1 {
		return "", nil
	}

	return matches[1], nil
}
