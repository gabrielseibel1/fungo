package unzip

import "github.com/gabrielseibel1/fungo/types"

func Map[T comparable, U any](m map[T]U) ([]T, []U) {
	t, u, i := make([]T, len(m)), make([]U, len(m)), 0
	for k, v := range m {
		t[i], u[i], i = k, v, i+1
	}
	return t, u
}

func Pairs[T any, U any](p []types.Pair[T, U]) ([]T, []U) {
	t, u := make([]T, len(p)), make([]U, len(p))
	for i, e := range p {
		t[i], u[i] = e.K, e.V
	}
	return t, u
}
