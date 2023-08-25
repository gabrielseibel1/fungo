package fungomap

func SomeK[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return keyInspectorOf[T, U](true)(m, f)
}

func NoK[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return keyInspectorOf[T, U](false)(m, f)
}

func SomeV[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return valInspectorOf[T, U](true)(m, f)
}

func NoV[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return valInspectorOf[T, U](false)(m, f)
}

func valInspectorOf[T comparable, U any](b bool) func(m map[T]U, f func(U) bool) bool {
	return func(m map[T]U, f func(U) bool) bool {
		for _, v := range m {
			if f(v) {
				return b
			}
		}
		return !b
	}
}

func keyInspectorOf[T comparable, U any](b bool) func(m map[T]U, f func(T) bool) bool {
	return func(m map[T]U, f func(T) bool) bool {
		for k := range m {
			if f(k) {
				return b
			}
		}
		return !b
	}
}
