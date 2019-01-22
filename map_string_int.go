package maps

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
