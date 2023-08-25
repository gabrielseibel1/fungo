package fungomap

type Pair[K, V any] struct {
	K K
	V V
}

func ToSliceK[T comparable, U any](m map[T]U) []T {
	r := make([]T, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func ToSliceV[T comparable, U any](m map[T]U) []U {
	r := make([]U, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func ToSliceP[T comparable, U any](m map[T]U) []Pair[T, U] {
	r := make([]Pair[T, U], 0, len(m))
	for k, v := range m {
		r = append(r, Pair[T, U]{k, v})
	}
	return r
}
