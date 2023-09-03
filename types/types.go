// Package types provides useful types that are used by other packages in the module
package types

// Pair is a type that contains two values of different generic types
type Pair[K, V any] struct {
	K K
	V V
}
