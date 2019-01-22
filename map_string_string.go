package maps

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
}
