package maps

func ContainsKeyStringBool(m map[string]bool, k string) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueStringBool(m map[string]bool, v bool) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysStringBool(m map[string]bool) []string {
	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesStringBool(m map[string]bool) []bool {
	values := []bool{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyStringBool(m map[string]bool) map[string]bool {
	copyMap := map[string]bool{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
