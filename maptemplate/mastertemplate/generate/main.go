package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	masterFilePath = "mapstringint/map_string_int.go"
)

func main() {
	masterTemplateContent := getMasterTemplateContent()
	err := ioutil.WriteFile("mastertemplate.go", []byte(masterTemplateContent), 0644)
	if err != nil {
		panic(err)
	}
}

func getMasterTemplateContent() string {
	masterTemplateString := getMasterTemplateString()
	masterTemplate := "package mastertemplate\n\nconst (\n\t"
	masterTemplate += fmt.Sprintf("String = `%s`", masterTemplateString)
	return masterTemplate
}

func getMasterTemplateString() string {
	masterFileString := getMasterFileString()
	replacements := map[string]string{
		"mastertemplate": "{{.Package}}",
		"String":         "{{.KeySuffix}}",
		"Int":            "{{.ValueSuffix}}",
		"string":         "{{.KeyType}}",
		"int":            "{{.ValueType}}",
	}
	replacedFileString := applyReplacements(masterFileString, replacements)
	return replacedFileString
}

func getMasterFileString() string {
	b, err := ioutil.ReadFile(masterFilePath)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func applyReplacements(original string, replacements map[string]string) string {
	replaced := original

	for old, new := range replacements {
		replaced = strings.Replace(replaced, old, new, -1)
	}

	return replaced
}
