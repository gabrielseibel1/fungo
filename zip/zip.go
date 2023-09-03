// Package zip has functions to zip pairs of slices into slices of pairs or maps
package zip

import "github.com/gabrielseibel1/fungo/types"

// SlicesToPairs zips two slices into a pairs slice of the size of the smallest slice provided
func SlicesToPairs[T any, U any](t []T, u []U) []types.Pair[T, U] {
	if len(t) <= len(u) {
		r := make([]types.Pair[T, U], len(t))
		for i := range t {
			r[i] = types.Pair[T, U]{K: t[i], V: u[i]}
		}
		return r
	} else {
		r := make([]types.Pair[T, U], len(u))
		for i := range u {
			r[i] = types.Pair[T, U]{K: t[i], V: u[i]}
		}
		return r
	}
}

// SlicesToMap zips two slices into a map of the size of the smallest slice provided
func SlicesToMap[T comparable, U any](t []T, u []U) map[T]U {
	r := make(map[T]U, len(t))
	for i := range t {
		r[t[i]] = u[i]
	}
	return r
}
