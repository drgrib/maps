package maptemplate

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type MapFile struct {
	Name        string
	MapTemplate MapTemplate
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
