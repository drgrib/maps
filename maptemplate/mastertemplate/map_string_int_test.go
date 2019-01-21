package mastertemplate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainsKeyStringInt_WhenContains_ReturnsTrue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := ContainsKeyStringInt(m, "b")

	assert.True(t, result)
}

func Test_ContainsKeyStringInt_WhenDoesNotContainStringValue_ReturnsFalse(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := ContainsKeyStringInt(m, "sandwich")

	assert.False(t, result)
}

func Test_ContainsKeyStringInt_WhenMapEmpty_ReturnsFalse(t *testing.T) {
	m := map[string]int{}

	result := ContainsKeyStringInt(m, "a")

	assert.False(t, result)
}
