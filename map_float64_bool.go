package maps

func ContainsKeyFloat64Bool(m map[float64]bool, k float64) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueFloat64Bool(m map[float64]bool, v bool) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysFloat64Bool(m map[float64]bool) []float64 {
	keys := []float64{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesFloat64Bool(m map[float64]bool) []bool {
	values := []bool{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyFloat64Bool(m map[float64]bool) map[float64]bool {
	copyMap := map[float64]bool{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
