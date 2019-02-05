package maps

func ContainsKeyIntFloat64(m map[int]float64, k int) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueIntFloat64(m map[int]float64, v float64) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysIntFloat64(m map[int]float64) []int {
	keys := []int{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesIntFloat64(m map[int]float64) []float64 {
	values := []float64{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyIntFloat64(m map[int]float64) map[int]float64 {
	copyMap := map[int]float64{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
