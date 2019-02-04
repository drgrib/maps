[![Build Status](https://travis-ci.com/drgrib/maps.svg?branch=master)](https://travis-ci.com/drgrib/maps)

# Installation

``` bash
go get https://github.com/drgrib/maps
```

# Purpose
This package provides type-safe implementations of "missing" `map` functions:

- `ContainsKey` - check if map contains key
- `ContainsValue` - check if map contains value
- `GetKeys` - get keys of a map
- `GetValues` - get values of a map
- `Copy` - deep copy of a map

implemented with these patterns

- `ContainsKeyKV(map[ktype]vtype, k ktype) bool`
- `ContainsValueKV(map[ktype]vtype, v vtype) bool`
- `GetKeysKV(map[ktype]vtype) []ktype`
- `GetValuesKV(map[ktype]vtype) []vtype`
- `CopyKV(map[ktype]vtype) map[ktype]vtype`

where `K` and `V` are the key and value for common map types:

- `StringString`
- `StringInt`
- `StringFloat32`
- `StringFloat64`
- `Float32Bool`
- `Float64Bool`
- `IntBool`
- `StringBool`

# Generation
The package extends beyond implementing these functions only for these basic type pairings. You can generate them for any `map` type, including maps with keys and values that are pointers, custom types, or `interface{}`.

There are two ways to do this. One is by installing and calling the `mapper` tool, which can be considered a sort of sibling to the [`stringer` tool](https://godoc.org/golang.org/x/tools/cmd/stringer). The other is programmatically by using the `maptemplate` library called within the [source code of `mapper`](https://github.com/drgrib/maps/blob/master/cmd/mapper/mapper.go) and [the code used to generate](https://github.com/drgrib/maps/blob/master/generate/main.go) the basic type `map` functions in this package.

## `mapper`

`mapper` can be installed using

``` bash
go get -u github.com/drgrib/maps/cmd/mapper
```

Then used with the `go generate` like this:

``` go
//go:generate mapper -types string:CustomType
```

Or directly on the commandline with the same command:

``` bash
mapper -types string:CustomType
```

Which, will generate the file `map_string_customvaluetype.go` that infers its `package` from surrounding `.go` files or the current folder name if no other files are found:

``` go
package maps

func ContainsKeyStringCustomValueType(m map[string]CustomType, k string) bool {
	_, ok := m[k]
	return ok
}

func ContainsValueStringCustomValueType(m map[string]CustomType, v CustomType) bool {
	for _, mValue := range m {
		if mValue == v {
			return true
		}
	}

	return false
}

func GetKeysStringCustomValueType(m map[string]CustomType) []string {
	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValuesStringCustomValueType(m map[string]CustomType) []CustomType {
	values := []CustomType{}

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func CopyStringCustomValueType(m map[string]CustomType) map[string]CustomType {
	copyMap := map[string]CustomType{}

	for k, v := range m {
		copyMap[k] = v
	}

	return copyMap
}
```

# About Generics
If generics are implemented for Go 2.0, these functions can be covered by a single file, without the need for the `mapper` tool and its underlying generation packages. Until then, there is this package.
