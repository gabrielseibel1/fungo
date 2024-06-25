// Package check has functions to check if some condition is valid for slices and maps
// TODO indexed
package check

// All checks that the provided slice has every element satifying the checker function
func All[T any](s []T, f func(T) bool) bool {
	return len(s) != 0 && !hasAny[T](false)(s, f)
}

// Some checks that the provided slice has any element that satisfies the checker function
func Some[T any](s []T, f func(T) bool) bool {
	return len(s) != 0 && hasAny[T](true)(s, f)
}

// None checks that the provided slice has no element that satisfies the checker function
func None[T any](s []T, f func(T) bool) bool {
	return !hasAny[T](true)(s, f)
}

// All checks that the provided map has every key satifying the checker function
func AllKeys[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return len(m) != 0 && !hasAnyKey[T, U](false)(m, f)
}

// SomeKey checks that the provided map has any key that satisfies the checker function
func SomeKey[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return hasAnyKey[T, U](true)(m, f)
}

// NoKey checks that the provided map has no key that satisfies the checker function
func NoKey[T comparable, U any](m map[T]U, f func(T) bool) bool {
	return !hasAnyKey[T, U](true)(m, f)
}

// All checks that the provided map has every value satifying the checker function
func AllValues[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return len(m) != 0 && !hasAnyVal[T, U](false)(m, f)
}

// SomeValue checks that the provided map has any value that satisfies the checker function
func SomeValue[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return hasAnyVal[T, U](true)(m, f)
}

// NoValue checks that the provided map has no value that satisfies the checker function
func NoValue[T comparable, U any](m map[T]U, f func(U) bool) bool {
	return !hasAnyVal[T, U](true)(m, f)
}

func hasAny[T any](b bool) func(s []T, f func(T) bool) bool {
	return func(s []T, f func(T) bool) bool {
		for _, v := range s {
			if f(v) == b {
				return true
			}
		}
		return false
	}
}

func hasAnyVal[T comparable, U any](b bool) func(m map[T]U, f func(U) bool) bool {
	return func(m map[T]U, f func(U) bool) bool {
		for _, v := range m {
			if f(v) == b {
				return true
			}
		}
		return false
	}
}

func hasAnyKey[T comparable, U any](b bool) func(m map[T]U, f func(T) bool) bool {
	return func(m map[T]U, f func(T) bool) bool {
		for k := range m {
			if f(k) == b {
				return true
			}
		}
		return false
	}
}
