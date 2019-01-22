package maptemplate

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func Test_Parse_WhenStringInt_MatchesExpected(t *testing.T) {
	mapTemplate := NewMapTemplate("maptemplatetest", "string", "int")

	content, err := mapTemplate.Parse()
	assert.Nil(t, err)
	assert.Equal(t, stripWhiteSpace(stringIntParsed), stripWhiteSpace(content))
}

func stripWhiteSpace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

const (
	stringIntParsed = `package maptemplatetest

	func ContainsKeyStringInt(m map[string]int, k string) bool {
			_, ok := m[k]
			return ok
	}

	func ContainsValueStringInt(m map[string]int, v int) bool {
			for _, mValue := range m {
					if mValue == v {
							return true
					}
			}

			return false
	}

	func GetKeysStringInt(m map[string]int) []string {
			keys := []string{}

			for k, _ := range m {
					keys = append(keys, k)
			}

			return keys
	}

	func GetValuesStringInt(m map[string]int) []int {
			values := []int{}

			for _, v := range m {
					values = append(values, v)
			}

			return values
	}

	func CopyStringInt(m map[string]int) map[string]int {
			copyMap := map[string]int{}

			for k, v := range m {
					copyMap[k] = v
			}

			return copyMap
	}
`
)
