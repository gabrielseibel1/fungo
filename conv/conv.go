package conv

func SliceToMapWithIndices[T any](s []T) map[int]T {
	r := make(map[int]T)
	for i, e := range s {
		r[i] = e
	}
	return r
}

func SliceToMapKeys[T comparable, U any](s []T, v func(T) U) map[T]U {
	r := make(map[T]U)
	for _, e := range s {
		r[e] = v(e)
	}
	return r
}

func SliceToMapValues[T any, U comparable](s []T, k func(T) U) map[U]T {
	r := make(map[U]T)
	for _, e := range s {
		r[k(e)] = e
	}
	return r
}

type Pair[K, V any] struct {
	K K
	V V
}

func MapKeysToSlice[T comparable, U any](m map[T]U) []T {
	r := make([]T, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func MapValuesToSlice[T comparable, U any](m map[T]U) []U {
	r := make([]U, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func MapToPairs[T comparable, U any](m map[T]U) []Pair[T, U] {
	r := make([]Pair[T, U], 0, len(m))
	for k, v := range m {
		r = append(r, Pair[T, U]{k, v})
	}
	return r
}

// TODO conv pairs to map
