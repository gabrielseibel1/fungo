// Package apply has functions to apply transformations to slices, maps, and channels
// TODO indexed
package apply

import "github.com/gabrielseibel1/fungo/types"

// Unindexed

// ToSlice applies a transforming function to a slice's elements, returning a slice with the return type of the function
func ToSlice[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i := range r {
		r[i] = f(s[i])
	}
	return r
}

// ToChannel creates a channel that has elements of the provided one, but with a transforming function applied.
// Does not consume the original channel, operates "on demand".
func ToChannel[T any, U any](c <-chan T, f func(T) U) chan U {
	r := make(chan U)
	go func(in <-chan T, out chan<- U) {
		for e := range in {
			out <- f(e)
		}
		close(out)
	}(c, r)
	return r
}

// ToKeys applies a transforming function to a map's keys, returning a map with keys of the return type of the function
func ToKeys[T comparable, U comparable, V any](m map[T]V, f func(T) U) map[U]V {
	r := make(map[U]V)
	for k, v := range m {
		r[f(k)] = v
	}
	return r
}

// ToValues applies a transforming function to a map's values, returning a map with values of the return type of the function
func ToValues[T comparable, U any, V any](m map[T]U, f func(U) V) map[T]V {
	r := make(map[T]V)
	for k, v := range m {
		r[k] = f(v)
	}
	return r
}

// ToPairs applies a transforming function to a slice of pairs, returning another with the return type of the function
func ToPairs[T any, U any, V any, W any](p []types.Pair[T, U], f func(types.Pair[T, U]) types.Pair[V, W]) []types.Pair[V, W] {
	q := make([]types.Pair[V, W], len(p))
	for i := range p {
		q[i] = f(p[i])
	}
	return q
}

// Indexed

// ToSliceIndexed is like ToSlice but with the index passed to the application function
func ToSliceIndexed[T, U any](s []T, f func(int, T) U) []U {
	r := make([]U, len(s))
	for i := range r {
		r[i] = f(i, s[i])
	}
	return r
}

// ToChannelIndexed is like ToChannel but with the index passed to the application function
func ToChannelIndexed[T any, U any](c <-chan T, f func(int, T) U) chan U {
	r := make(chan U)
	var i int
	go func(in <-chan T, out chan<- U) {
		for e := range in {
			out <- f(i, e)
			i++
		}
		close(out)
	}(c, r)
	return r
}

// ToKeysValued is like ToKeys but with the values passed to the application function
func ToKeysValued[T comparable, U comparable, V any](m map[T]V, f func(T, V) U) map[U]V {
	r := make(map[U]V)
	for k, v := range m {
		r[f(k, v)] = v
	}
	return r
}

// ToValueKeyed is like ToValues but with the keys passed to the application function
func ToValuesKeyed[T comparable, U any, V any](m map[T]U, f func(T, U) V) map[T]V {
	r := make(map[T]V)
	for k, v := range m {
		r[k] = f(k, v)
	}
	return r
}

// ToPairsIndexed is like ToPairs but with the indexed passed to the application function
func ToPairsIndexed[T any, U any, V any, W any](p []types.Pair[V, U], f func(int, types.Pair[V, U]) types.Pair[V, W]) []types.Pair[V, W] {
	q := make([]types.Pair[V, W], len(p))
	for i := range p {
		q[i] = f(i, p[i])
	}
	return q
}
