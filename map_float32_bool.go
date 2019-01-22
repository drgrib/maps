package maps

func ContainsKeyFloat32Bool(m map[float32]bool, k float32) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueFloat32Bool(m map[float32]bool, v bool) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysFloat32Bool(m map[float32]bool) []float32 {
	keys := []float32{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesFloat32Bool(m map[float32]bool) []bool {
	values := []bool{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyFloat32Bool(m map[float32]bool) map[float32]bool {
	copyMap := map[float32]bool{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
