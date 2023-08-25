package fungomap

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
