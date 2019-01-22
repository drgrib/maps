package mastertemplate

const (
	String = `package {{.Package}}

func ContainsKey{{.KeySuffix}}{{.ValueSuffix}}(m map[{{.KeyType}}]{{.ValueType}}, k {{.KeyType}}) bool {
	_, ok := m[k]
	return ok
}

func ContainsValue{{.KeySuffix}}{{.ValueSuffix}}(m map[{{.KeyType}}]{{.ValueType}}, v {{.ValueType}}) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeys{{.KeySuffix}}{{.ValueSuffix}}(m map[{{.KeyType}}]{{.ValueType}}) []{{.KeyType}} {
	keys := []{{.KeyType}}{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValues{{.KeySuffix}}{{.ValueSuffix}}(m map[{{.KeyType}}]{{.ValueType}}) []{{.ValueType}} {
	values := []{{.ValueType}}{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func Copy{{.KeySuffix}}{{.ValueSuffix}}(m map[{{.KeyType}}]{{.ValueType}}) map[{{.KeyType}}]{{.ValueType}} {
	copyMap := map[{{.KeyType}}]{{.ValueType}}{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
`
)