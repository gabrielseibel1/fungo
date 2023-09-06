// Package filter has functions to filter slices, maps, and channels by some condition
package filter

// Slice filters a slice by a function that takes an element an returns a bool,
// returning another slice with only the elements that satisfy the function condition
func Slice[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// SliceIndexed filters a slice by a function that takes an index and an element and returns a bool,
// returning another slice with only the elements that satisfy the function condition
func SliceIndexed[T any](s []T, f func(int, T) bool) []T {
	r := make([]T, 0)
	for i, v := range s {
		if f(i, v) {
			r = append(r, v)
		}
	}
	return r
}

// Channel filters a channel by a function, returning another channel with only the elements that satisfy the function condition
// Does not consume the whole channel, works "on demand"
func Channel[T any](c <-chan T, f func(T) bool) chan T {
	r := make(chan T)
	go func(in <-chan T, out chan<- T) {
		for e := range in {
			if f(e) {
				out <- e
			}
		}
		close(out)
	}(c, r)
	return r
}

// MapByKeys filters a map by a function, returning another map with only the keys satisfy the function condition
// TODO just "Map", with keys and values passed to func
func MapByKeys[T comparable, U any](m map[T]U, f func(T) bool) map[T]U {
	r := make(map[T]U)
	for k, v := range m {
		if f(k) {
			r[k] = v
		}
	}
	return r
}

// MapByValues filters a map by a function, returning another map with only the values satisfy the function condition
func MapByValues[T comparable, U any](m map[T]U, f func(U) bool) map[T]U {
	r := make(map[T]U)
	for k, v := range m {
		if f(v) {
			r[k] = v
		}
	}
	return r
}
