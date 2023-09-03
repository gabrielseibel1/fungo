// Package fold has functions to aggregate/reduce elements in slices
// TODO indexed
// TODO map
package fold

// Slice incrementally combines a seed value and then values returned by a function applied to the slice elements,
// reducing the slice down to a single value, which is returned
func Slice[T any, U any](s []T, i U, f func(U, T) U) U {
	r := i
	for _, e := range s {
		r = f(r, e)
	}
	return r
}
