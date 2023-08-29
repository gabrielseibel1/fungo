package conv

import "github.com/gabrielseibel1/fungo/types"

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

func MapToPairs[T comparable, U any](m map[T]U) []types.Pair[T, U] {
	r := make([]types.Pair[T, U], 0, len(m))
	for k, v := range m {
		r = append(r, types.Pair[T, U]{K: k, V: v})
	}
	return r
}

func PairsToMap[T comparable, U any](p []types.Pair[T, U]) map[T]U {
	r := make(map[T]U)
	for _, e := range p {
		r[e.K] = e.V
	}
	return r
}

func PairsKeysToSlice[T any, U any](p []types.Pair[T, U]) []T {
	r := make([]T, len(p))
	for i, e := range p {
		r[i] = e.K
	}
	return r
}

func PairsValuesToSlice[T any, U any](p []types.Pair[T, U]) []U {
	r := make([]U, len(p))
	for i, e := range p {
		r[i] = e.V
	}
	return r
}
