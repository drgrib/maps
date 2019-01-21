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

func Test_ContainsValueStringInt_WhenContains_ReturnsTrue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := ContainsValueStringInt(m, 3)

	assert.True(t, result)
}

func Test_ContainsValueStringInt_WhenDoesNotContainStringValue_ReturnsFalse(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := ContainsValueStringInt(m, 42)

	assert.False(t, result)
}

func Test_ContainsValueStringInt_WhenMapEmpty_ReturnsFalse(t *testing.T) {
	m := map[string]int{}

	result := ContainsValueStringInt(m, 1)

	assert.False(t, result)
}

func Test_GetKeysStringInt_WhenKeys_ReturnsKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	values := GetKeysStringInt(m)

	expected := []string{"a", "b", "c"}

	assert.Equal(t, len(expected), len(values))

	for _, v := range expected {
		assert.Contains(t, values, v)
	}

	for _, v := range values {
		assert.Contains(t, expected, v)
	}
}

func Test_GetKeysStringInt_WhenEmpty_ReturnsEmpty(t *testing.T) {
	m := map[string]int{}

	values := GetKeysStringInt(m)

	assert.Equal(t, []string{}, values)
}

func Test_GetValuesStringInt_WhenValues_ReturnsValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	values := GetValuesStringInt(m)

	expected := []int{1, 2, 3}

	assert.Equal(t, len(expected), len(values))

	for _, v := range expected {
		assert.Contains(t, values, v)
	}

	for _, v := range values {
		assert.Contains(t, expected, v)
	}
}

func Test_GetValuesStringInt_WhenEmpty_ReturnsEmpty(t *testing.T) {
	m := map[string]int{}

	values := GetValuesStringInt(m)

	assert.Equal(t, []int{}, values)
}

func Test_CopyStringInt_WhenValues_ReturnsDeepCopy(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	deepCopyMap := CopyStringInt(m)

	assert.Equal(t, len(m), len(deepCopyMap))

	for k, v := range m {
		checkValue, ok := deepCopyMap[k]
		assert.True(t, ok)
		assert.Equal(t, v, checkValue)
	}

	for k, v := range deepCopyMap {
		checkValue, ok := m[k]
		assert.True(t, ok)
		assert.Equal(t, v, checkValue)
	}
}

func Test_CopyStringInt_WhenEmpty_ReturnsEmptyDeepCopy(t *testing.T) {
	m := map[string]int{}

	deepCopyMap := CopyStringInt(m)

	assert.Equal(t, len(m), len(deepCopyMap))
	assert.Equal(t, len(m), 0)
	assert.Equal(t, len(deepCopyMap), 0)
}
