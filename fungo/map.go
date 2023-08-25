package fungo

func Mapped[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i := range r {
		r[i] = f(s[i])
	}
	return r
}

func MappedK[T comparable, U comparable, V any](m map[T]V, f func(T) U) map[U]V {
	r := make(map[U]V)
	for k, v := range m {
		r[f(k)] = v
	}
	return r
}

func MappedV[T comparable, U any, V any](m map[T]U, f func(U) V) map[T]V {
	r := make(map[T]V)
	for k, v := range m {
		r[k] = f(v)
	}
	return r
}
