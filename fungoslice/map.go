package fungoslice

func Mapped[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i := range r {
		r[i] = f(s[i])
	}
	return r
}
