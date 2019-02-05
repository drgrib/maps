package maps

func ContainsKeyIntFloat32(m map[int]float32, k int) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueIntFloat32(m map[int]float32, v float32) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysIntFloat32(m map[int]float32) []int {
	keys := []int{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesIntFloat32(m map[int]float32) []float32 {
	values := []float32{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyIntFloat32(m map[int]float32) map[int]float32 {
	copyMap := map[int]float32{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
