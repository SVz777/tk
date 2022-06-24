package collections

type Set[T comparable] map[T]Empty

func NewSet[T comparable](keys ...T) Set[T] {
	s := make(Set[T], len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}

func (s Set[T]) AllItems() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) IsContain(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Add(key T) {
	s[key] = Empty{}
}

func (s Set[T]) Delete(key T) {
	delete(s, key)
}
