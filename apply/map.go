// Package apply has functions to apply transformations to slices, maps, and channels
// TODO indexed
package apply

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
