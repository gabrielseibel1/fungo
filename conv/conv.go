// Package conv has functions to convert between slices and maps
// TODO indexed
package conv

import "github.com/gabrielseibel1/fungo/types"

// SliceToMapWithIndices converts a slice to a map, with keys as the slice's indices, and values as the slice's elements
func SliceToMapWithIndices[T any](s []T) map[int]T {
	r := make(map[int]T)
	for i, e := range s {
		r[i] = e
	}
	return r
}

// SliceToMapKeys converts a slice to a map, with keys as the slice's elements, and values as the slice's elements transformed by a function
func SliceToMapKeys[T comparable, U any](s []T, v func(T) U) map[T]U {
	r := make(map[T]U)
	for _, e := range s {
		r[e] = v(e)
	}
	return r
}

// SliceToMapValues converts a slice to a map, with values as the slice's elements, and key as the slice's elements transformed by a function
func SliceToMapValues[T any, U comparable](s []T, k func(T) U) map[U]T {
	r := make(map[U]T)
	for _, e := range s {
		r[k(e)] = e
	}
	return r
}

// MapKeysToSlice converts a map to a slice, with the elements as the map keys
func MapKeysToSlice[T comparable, U any](m map[T]U) []T {
	r := make([]T, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// MapValuesToSlice converts a map to a slice, with the elements as the map values
func MapValuesToSlice[T comparable, U any](m map[T]U) []U {
	r := make([]U, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// MapToPairs converts a map to a slice of pairs, with the elements as the map keys and values grouped in pairs
func MapToPairs[T comparable, U any](m map[T]U) []types.Pair[T, U] {
	r := make([]types.Pair[T, U], 0, len(m))
	for k, v := range m {
		r = append(r, types.Pair[T, U]{K: k, V: v})
	}
	return r
}

// PairsToMap converts a pairs slice to a map, mapping each pair's key to the respective value
func PairsToMap[T comparable, U any](p []types.Pair[T, U]) map[T]U {
	r := make(map[T]U)
	for _, e := range p {
		r[e.K] = e.V
	}
	return r
}

// PairsKeysToSlice converts a pairs' slice to a slice of the pairs' keys
func PairsKeysToSlice[T any, U any](p []types.Pair[T, U]) []T {
	r := make([]T, len(p))
	for i, e := range p {
		r[i] = e.K
	}
	return r
}

// PairsValuesToSlice converts a pairs' slice to a slice of the pairs' values
func PairsValuesToSlice[T any, U any](p []types.Pair[T, U]) []U {
	r := make([]U, len(p))
	for i, e := range p {
		r[i] = e.V
	}
	return r
}
