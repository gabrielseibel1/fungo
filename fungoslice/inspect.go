package fungoslice

func Some[T any](s []T, f func(T) bool) bool {
	return inspectorOf[T](true)(s, f)
}

func None[T any](s []T, f func(T) bool) bool {
	return inspectorOf[T](false)(s, f)
}

type inspector[T any] func(s []T, f func(T) bool) bool

func inspectorOf[T any](b bool) inspector[T] {
	return func(s []T, f func(T) bool) bool {
		for _, v := range s {
			if f(v) {
				return b
			}
		}
		return !b
	}
}
