package types

// Pair is a type that contains two values of different generic types
type Pair[K, V any] struct {
	K K
	V V
}
