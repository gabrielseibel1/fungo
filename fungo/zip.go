package fungo

func Zip[T any, U any](t []T, u []U) []Pair[T, U] {
	if len(t) <= len(u) {
		r := make([]Pair[T, U], len(t))
		for i := range t {
			r[i] = Pair[T, U]{t[i], u[i]}
		}
		return r
	} else {
		r := make([]Pair[T, U], len(u))
		for i := range u {
			r[i] = Pair[T, U]{t[i], u[i]}
		}
		return r
	}
}
