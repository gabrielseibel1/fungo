package fungo

func Filtered[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
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
