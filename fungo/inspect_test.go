package fungo

import (
	"testing"
)

func TestSome(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[Data]{
		{
			name: "empty/nil",
			args: args[Data]{
				s: nil,
				f: PassAll[Data],
			},
			want: false,
		},
		{
			name: "some",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: Is(Data1),
			},
			want: true,
		},
		{
			name: "none",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: PassNo[Data],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[Data]{
		{
			name: "empty/nil",
			args: args[Data]{
				s: nil,
				f: PassAll[Data],
			},
			want: true,
		},
		{
			name: "some",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: Is(Data1),
			},
			want: false,
		},
		{
			name: "none",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: PassNo[Data],
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := None(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoK(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
		f func(T) bool
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "empty",
			args: args[string, int]{
				m: nil,
				f: PassAll[string],
			},
			want: true,
		},
		{
			name: "some",
			args: args[string, int]{
				m: map[string]int{"1": 1, "b": 2},
				f: IsNumber,
			},
			want: false,
		},
		{
			name: "none",
			args: args[string, int]{
				m: map[string]int{"a": 1, "b": 2},
				f: IsNumber,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoK(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("NoK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoV(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
		f func(U) bool
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want bool
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				m: nil,
				f: PassAll[string],
			},
			want: true,
		},
		{
			name: "some",
			args: args[int, string]{
				m: map[int]string{1: "1", 2: "b"},
				f: IsNumber,
			},
			want: false,
		},
		{
			name: "nome",
			args: args[int, string]{
				m: map[int]string{1: "a", 2: "b"},
				f: IsNumber,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoV(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("NoV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeK(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
		f func(T) bool
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "empty",
			args: args[string, int]{
				m: nil,
				f: PassAll[string],
			},
			want: false,
		},
		{
			name: "some",
			args: args[string, int]{
				m: map[string]int{"1": 1, "b": 2},
				f: IsNumber,
			},
			want: true,
		},
		{
			name: "none",
			args: args[string, int]{
				m: map[string]int{"a": 1, "b": 2},
				f: IsNumber,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SomeK(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("SomeK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeV(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
		f func(U) bool
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want bool
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				m: nil,
				f: PassAll[string],
			},
			want: false,
		},
		{
			name: "some",
			args: args[int, string]{
				m: map[int]string{1: "1", 2: "b"},
				f: IsNumber,
			},
			want: true,
		},
		{
			name: "nome",
			args: args[int, string]{
				m: map[int]string{1: "a", 2: "b"},
				f: IsNumber,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SomeV(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("SomeV() = %v, want %v", got, tt.want)
			}
		})
	}
}
