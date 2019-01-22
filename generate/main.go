package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/drgrib/maps/maptemplate"
)

func main() {
	pairMap := map[string]bool{}
	addStringKeyPairs(pairMap)
	addBoolValuePairs(pairMap)
	writePairs(pairMap)
}

func addStringKeyPairs(pairMap map[string]bool) {
	keys := []string{"string"}
	values := []string{
		"string",
		"int",
		"float32",
		"float64",
		"bool",
	}

	addPairs(keys, values, pairMap)
}

func addBoolValuePairs(pairMap map[string]bool) {
	keys := []string{
		"string",
		"int",
		"float32",
		"float64",
	}
	values := []string{
		"bool",
	}

	addPairs(keys, values, pairMap)
}

func writePairs(pairMap map[string]bool) {
	for pair := range pairMap {
		pairSplit := strings.Split(pair, ":")
		key := pairSplit[0]
		value := pairSplit[1]
		writePair(key, value)
	}
}

func writePair(key, value string) {
	mapTemplate := maptemplate.NewMapTemplate("maps", key, value)
	writeMapTemplate(mapTemplate)
}

func writeMapTemplate(mapTemplate maptemplate.MapTemplate) {
	fileName := fmt.Sprintf("map_%s_%s.go", strings.ToLower(mapTemplate.KeySuffix), strings.ToLower(mapTemplate.ValueSuffix))

	content, err := mapTemplate.Parse()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func addPairs(keys, values []string, pairMap map[string]bool) {
	for _, k := range keys {
		for _, v := range values {
			pair := fmt.Sprintf("%s:%s", k, v)
			pairMap[pair] = true
		}
	}
}
