package main // import "github.com/drgrib/maps/cmd/mapper"

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/drgrib/maps/maptemplate"
)

func main() {
	types := flag.String("types", "", "key:value type pair (required)")
	suffix := flag.String("suffix", "", "custom function suffix (optional)")
	output := flag.String("output", "", "target file name (optional)")
	packageName := flag.String("package", "", "target package name (optional)")
	flag.Parse()

	if *types == "" {
		exitWithError("Error: At least one key:value -types pair required")
	}

	mapFile, err := parseToDefaultMapFile(*types)
	if err != nil {
		exitWithError(err.Error())
	}

	if *suffix != "" {
		mapFile.MapTemplate.KeySuffix = *suffix
		mapFile.MapTemplate.ValueSuffix = ""
	}

	if *output != "" {
		mapFile.FileName = *output
	}

	if *packageName != "" {
		mapFile.MapTemplate.Package = *packageName
	}

	err = mapFile.Write()
	if err != nil {
		exitWithError(err.Error())
	}
}

func parseToDefaultMapFile(types string) (*maptemplate.MapFile, error) {
	pairSplit := strings.Split(types, ":")
	if len(pairSplit) != 2 {
		return nil, fmt.Errorf("Error: Bad types string %q", types)
	}
	keytype, valuetype := pairSplit[0], pairSplit[1]

	mapFile, err := maptemplate.NewMapFile(keytype, valuetype)
	if err != nil {
		return nil, err
	}

	return mapFile, nil
}

func exitWithError(message string) {
	fmt.Fprintf(os.Stderr, message+"\n")
	os.Exit(2)
}
