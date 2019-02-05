package maps

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
}
