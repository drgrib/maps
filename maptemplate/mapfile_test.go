package maptemplate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
