package collections

func Union[T comparable](arrs ...[]T) []T {
	if len(arrs) == 0 {
		return nil
	} else if len(arrs) == 1 {
		return arrs[0]
	}
	values := NewSet[T]()
	for _, arr := range arrs {
		for _, v := range arr {
			values.Add(v)
		}
	}
	return values.AllItems()
}

func Inter[T comparable](arrs ...[]T) []T {
	if len(arrs) == 0 {
		return nil
	} else if len(arrs) == 1 {
		return arrs[0]
	}
	values := NewSet(arrs[0]...)
	for _, arr := range arrs[1:] {
		tSet := NewSet(arr...)
		for v := range values {
			if !tSet.IsContain(v) {
				values.Delete(v)
			}
		}
	}
	return values.AllItems()
}

func InSlice[T comparable](v T, arr []T) bool {
	for _, vv := range arr {
		if vv == v {
			return true
		}
	}
	return false
}

func ToAnySlice[T comparable](arr []T) []any {
	ret := make([]any, len(arr))
	for idx, item := range arr {
		ret[idx] = item
	}
	return ret
}

func Map[T, R any](arr []T, f func(int, T) R) []R {
	ret := make([]R, len(arr))
	for idx, item := range arr {
		ret[idx] = f(idx, item)
	}
	return ret
}

func Reduce[T, R any](arr []T, f func(int, T, R) R) R {
	var ret R
	for idx, item := range arr {
		ret = f(idx, item, ret)
	}
	return ret
}

func Filter[T any](arr []T, f func(int, T) bool) []T {
	var ret []T
	for idx, item := range arr {
		if f(idx, item) {
			ret = append(ret, item)
		}
	}
	return ret
}
