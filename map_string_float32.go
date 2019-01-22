package maps

func ContainsKeyStringFloat32(m map[string]float32, k string) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueStringFloat32(m map[string]float32, v float32) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysStringFloat32(m map[string]float32) []string {
	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesStringFloat32(m map[string]float32) []float32 {
	values := []float32{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyStringFloat32(m map[string]float32) map[string]float32 {
	copyMap := map[string]float32{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
