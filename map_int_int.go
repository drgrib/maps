package maps

func ContainsKeyIntInt(m map[int]int, k int) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueIntInt(m map[int]int, v int) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysIntInt(m map[int]int) []int {
	keys := []int{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesIntInt(m map[int]int) []int {
	values := []int{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyIntInt(m map[int]int) map[int]int {
	copyMap := map[int]int{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
