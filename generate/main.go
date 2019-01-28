package main

import (
	"fmt"
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

func addPairs(keys, values []string, pairMap map[string]bool) {
	for _, k := range keys {
		for _, v := range values {
			pair := fmt.Sprintf("%s:%s", k, v)
			pairMap[pair] = true
		}
	}
}

func writePairs(pairMap map[string]bool) {
	for pair := range pairMap {
		pairSplit := strings.Split(pair, ":")
		key := pairSplit[0]
		value := pairSplit[1]

		mapFile, err := maptemplate.NewMapFile(key, value)
		if err != nil {
			panic(err)
		}

		err = mapFile.Write()
		if err != nil {
			panic(err)
		}
	}
}
