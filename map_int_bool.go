package maps

func ContainsKeyIntBool(m map[int]bool, k int) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueIntBool(m map[int]bool, v bool) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysIntBool(m map[int]bool) []int {
	keys := []int{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesIntBool(m map[int]bool) []bool {
	values := []bool{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyIntBool(m map[int]bool) map[int]bool {
	copyMap := map[int]bool{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
