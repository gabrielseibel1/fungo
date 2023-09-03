package check

// TODO indexed
// TODO all

// Some checks that the provided slice has any element that satisfies the checker function
func Some[T any](s []T, f func(T) bool) bool {
	return check[T](true)(s, f)
}

// None checks that the provided slice has no element that satisfies the checker function
func None[T any](s []T, f func(T) bool) bool {
	return check[T](false)(s, f)
}

// SomeKey checks that the provided map has any key that satisfies the checker function
func SomeKey[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return checkKey[T, U](true)(m, f)
}

// NoKey checks that the provided map has no key that satisfies the checker function
func NoKey[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return checkKey[T, U](false)(m, f)
}

// SomeValue checks that the provided map has any value that satisfies the checker function
func SomeValue[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return checkVal[T, U](true)(m, f)
}

// NoValue checks that the provided map has no value that satisfies the checker function
func NoValue[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return checkVal[T, U](false)(m, f)
}

func check[T any](b bool) func(s []T, f func(T) bool) bool {
	return func(s []T, f func(T) bool) bool {
		for _, v := range s {
			if f(v) {
				return b
			}
		}
		return !b
	}
}
func checkVal[T comparable, U any](b bool) func(m map[T]U, f func(U) bool) bool {
	return func(m map[T]U, f func(U) bool) bool {
		for _, v := range m {
			if f(v) {
				return b
			}
		}
		return !b
	}
}

func checkKey[T comparable, U any](b bool) func(m map[T]U, f func(T) bool) bool {
	return func(m map[T]U, f func(T) bool) bool {
		for k := range m {
			if f(k) {
				return b
			}
		}
		return !b
	}
}
