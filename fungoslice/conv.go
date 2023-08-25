package fungoslice

func ToMapI[T any](s []T) map[int]T {
	r := make(map[int]T)
	for i, e := range s {
		r[i] = e
	}
	return r
}

func ToMapK[T comparable, U any](s []T, v func(T) U) map[T]U {
	r := make(map[T]U)
	for _, e := range s {
		r[e] = v(e)
	}
	return r
}

func ToMapV[T any, U comparable](s []T, k func(T) U) map[U]T {
	r := make(map[U]T)
	for _, e := range s {
		r[k(e)] = e
	}
	return r
}
