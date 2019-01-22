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

func Test_Parse_WhenIntString_MatchesExpected(t *testing.T) {
	mapTemplate := NewMapTemplate("maptemplatetest", "int", "string")

	content, err := mapTemplate.Parse()

	assert.Nil(t, err)
	assert.Equal(t, stripWhiteSpace(intStringParsed), stripWhiteSpace(content))
}

func Test_Parse_WhenStringString_MatchesExpected(t *testing.T) {
	mapTemplate := NewMapTemplate("maptemplatetest", "string", "string")

	content, err := mapTemplate.Parse()

	assert.Nil(t, err)
	assert.Equal(t, stripWhiteSpace(stringStringParsed), stripWhiteSpace(content))
}

func Test_Parse_WhenStringCreatedtype_MatchesExpected(t *testing.T) {
	mapTemplate := NewMapTemplate("maptemplatetest", "string", "Createdtype")

	content, err := mapTemplate.Parse()

	assert.Nil(t, err)
	assert.Equal(t, stripWhiteSpace(stringCreatedtypeParsed), stripWhiteSpace(content))
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
}`

	intStringParsed = `package maptemplatetest

func ContainsKeyIntString(m map[int]string, k int) bool {
		_, ok := m[k]
		return ok
}

func ContainsValueIntString(m map[int]string, v string) bool {
		for _, mValue := range m {
				if mValue == v {
						return true
				}
		}

		return false
}

func GetKeysIntString(m map[int]string) []int {
		keys := []int{}

		for k, _ := range m {
				keys = append(keys, k)
		}

		return keys
}

func GetValuesIntString(m map[int]string) []string {
		values := []string{}

		for _, v := range m {
				values = append(values, v)
		}

		return values
}

func CopyIntString(m map[int]string) map[int]string {
		copyMap := map[int]string{}

		for k, v := range m {
				copyMap[k] = v
		}

		return copyMap
}`

	stringStringParsed = `package maptemplatetest

func ContainsKeyStringString(m map[string]string, k string) bool {
        _, ok := m[k]
        return ok
}

func ContainsValueStringString(m map[string]string, v string) bool {
        for _, mValue := range m {
                if mValue == v {
                        return true
                }
        }

        return false
}

func GetKeysStringString(m map[string]string) []string {
        keys := []string{}

        for k, _ := range m {
                keys = append(keys, k)
        }

        return keys
}

func GetValuesStringString(m map[string]string) []string {
        values := []string{}

        for _, v := range m {
                values = append(values, v)
        }

        return values
}

func CopyStringString(m map[string]string) map[string]string {
        copyMap := map[string]string{}

        for k, v := range m {
                copyMap[k] = v
        }

        return copyMap
}`

	stringCreatedtypeParsed = `package maptemplatetest

func ContainsKeyStringCreatedtype(m map[string]Createdtype, k string) bool {
        _, ok := m[k]
        return ok
}

func ContainsValueStringCreatedtype(m map[string]Createdtype, v Createdtype) bool {
        for _, mValue := range m {
                if mValue == v {
                        return true
                }
        }

        return false
}

func GetKeysStringCreatedtype(m map[string]Createdtype) []string {
        keys := []string{}

        for k, _ := range m {
                keys = append(keys, k)
        }

        return keys
}

func GetValuesStringCreatedtype(m map[string]Createdtype) []Createdtype {
        values := []Createdtype{}

        for _, v := range m {
                values = append(values, v)
        }

        return values
}

func CopyStringCreatedtype(m map[string]Createdtype) map[string]Createdtype {
        copyMap := map[string]Createdtype{}

        for k, v := range m {
                copyMap[k] = v
        }

        return copyMap
}`
)
