package maps

func ContainsKeyStringFloat64(m map[string]float64, k string) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueStringFloat64(m map[string]float64, v float64) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysStringFloat64(m map[string]float64) []string {
	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesStringFloat64(m map[string]float64) []float64 {
	values := []float64{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyStringFloat64(m map[string]float64) map[string]float64 {
	copyMap := map[string]float64{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
