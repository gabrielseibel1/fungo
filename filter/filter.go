package filter

func Slice[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Channel[T any](c <-chan T, f func(T) bool) chan T {
	r := make(chan T)
	go func(in <-chan T, out chan<- T) {
		for e := range in {
			if f(e) {
				out <- e
			}
		}
		close(out)
	}(c, r)
	return r
}

func MapByKeys[T comparable, U any](m map[T]U, f func(T) bool) map[T]U {
	r := make(map[T]U)
	for k, v := range m {
		if f(k) {
			r[k] = v
		}
	}
	return r
}

func MapByValues[T comparable, U any](m map[T]U, f func(U) bool) map[T]U {
	r := make(map[T]U)
	for k, v := range m {
		if f(v) {
			r[k] = v
		}
	}
	return r
}
