package fungo

func Fold[T any, U any](s []T, i U, f func(U, T) U) U {
	r := i
	for _, e := range s {
		r = f(r, e)
	}
	return r
}
