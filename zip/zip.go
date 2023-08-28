package zip

import "github.com/gabrielseibel1/fungo/conv"

func Slices[T any, U any](t []T, u []U) []conv.Pair[T, U] {
	if len(t) <= len(u) {
		r := make([]conv.Pair[T, U], len(t))
		for i := range t {
			r[i] = conv.Pair[T, U]{K: t[i], V: u[i]}
		}
		return r
	} else {
		r := make([]conv.Pair[T, U], len(u))
		for i := range u {
			r[i] = conv.Pair[T, U]{K: t[i], V: u[i]}
		}
		return r
	}
}

// TODO unzip
