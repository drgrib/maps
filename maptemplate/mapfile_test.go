package maptemplate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMapFile_WhenCalled_ReturnsCorrectMapFile(t *testing.T) {
	keytype := "string"
	valuetype := "*DerivedType"
	mapFile, err := NewMapFile(keytype, valuetype)
	assert.Nil(t, err)

	assert.Equal(t, "map_string_derivedtype.go", mapFile.FileName)
	assert.Equal(t, "maptemplate", mapFile.MapTemplate.Package)
	assert.Equal(t, keytype, mapFile.MapTemplate.KeyType)
	assert.Equal(t, valuetype, mapFile.MapTemplate.ValueType)
	assert.Equal(t, "String", mapFile.MapTemplate.KeySuffix)
	assert.Equal(t, "DerivedType", mapFile.MapTemplate.ValueSuffix)

	parsed, err := mapFile.MapTemplate.Parse()
	assert.Nil(t, err)
	assert.NotEmpty(t, parsed)
}

func Test_getCurrentDir_WhenCalled_ReturnsDir(t *testing.T) {
	dir, err := getCurrentDir()
	assert.Nil(t, err)

	assert.Equal(t, "maptemplate", dir)
}

func Test_getWdGoFiles_WhenCalled_ReturnsFiles(t *testing.T) {
	goFiles, err := getWdGoFiles()
	assert.Nil(t, err)

	assert.True(t, len(goFiles) > 0)
}

func Test_getFilePackage_WhenCalled_GetsPackage(t *testing.T) {
	packageName, err := getFilePackage("mapfile.go")
	assert.Nil(t, err)

	assert.Equal(t, "maptemplate", packageName)
}

func Test_inferPackage_WhenCalled_GetsPackage(t *testing.T) {
	packageName, err := inferPackage()
	assert.Nil(t, err)

	assert.Equal(t, "maptemplate", packageName)
}
