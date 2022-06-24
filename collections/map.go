package collections

func Keys[KT comparable, VT any](src map[KT]VT) []KT {
	keys := make([]KT, 0, len(src))
	for k := range src {
		keys = append(keys, k)
	}
	return keys
}

func Values[KT comparable, VT any](src map[KT]VT) []VT {
	values := make([]VT, 0, len(src))
	for _, v := range src {
		values = append(values, v)
	}
	return values
}

func Flip[KT, VT comparable](src map[KT]VT) map[VT]KT {
	newMap := make(map[VT]KT, len(src))
	for k, v := range src {
		newMap[v] = k
	}
	return newMap
}

func Copy[KT comparable, VT any](src map[KT]VT) map[KT]VT {
	newMap := make(map[KT]VT, len(src))
	for k, v := range src {
		newMap[k] = v
	}
	return newMap
}
