package collections

type Counter[T comparable] map[T]uint64

func NewCounter[T comparable](keys ...T) Counter[T] {
	s := make(Counter[T], len(keys))
	for _, key := range keys {
		s.AddOne(key)
	}
	return s
}

func (c Counter[T]) Count(key T) uint64 {
	return c[key]
}

func (c Counter[T]) AddOne(key T) {
	c.Add(key, 1)
}

func (c Counter[T]) Add(key T, v uint64) {
	if _, ok := c[key]; ok {
		c[key] += v
	} else {
		c[key] = v
	}
}

func (c Counter[T]) SubOne(key T) {
	c.Sub(key, 1)
}

func (c Counter[T]) Sub(key T, v uint64) {
	if vv, ok := c[key]; ok && vv > v {
		c[key] -= v
	} else {
		c[key] = 0
	}
}
