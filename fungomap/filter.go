package fungomap

func FilteredByK[T comparable, U any](m map[T]U, f func(T) bool) map[T]U {
	r := make(map[T]U)
	for k, v := range m {
		if f(k) {
			r[k] = v
		}
	}
	return r
}

func FilteredByV[T comparable, U any](m map[T]U, f func(U) bool) map[T]U {
	r := make(map[T]U)
	for k, v := range m {
		if f(v) {
			r[k] = v
		}
	}
	return r
}
