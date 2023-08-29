package check

import (
	"github.com/gabrielseibel1/fungo/util"
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
	tests := []testCase[util.Data]{
		{
			name: "empty/nil",
			args: args[util.Data]{
				s: nil,
				f: util.PassAll[util.Data],
			},
			want: false,
		},
		{
			name: "some",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.Is(util.Data1),
			},
			want: true,
		},
		{
			name: "none",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassNo[util.Data],
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
	tests := []testCase[util.Data]{
		{
			name: "empty/nil",
			args: args[util.Data]{
				s: nil,
				f: util.PassAll[util.Data],
			},
			want: true,
		},
		{
			name: "some",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.Is(util.Data1),
			},
			want: false,
		},
		{
			name: "none",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassNo[util.Data],
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

func TestNoKey(t *testing.T) {
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
				f: util.PassAll[string],
			},
			want: true,
		},
		{
			name: "some",
			args: args[string, int]{
				m: map[string]int{"1": 1, "b": 2},
				f: util.IsNumber,
			},
			want: false,
		},
		{
			name: "none",
			args: args[string, int]{
				m: map[string]int{"a": 1, "b": 2},
				f: util.IsNumber,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoKey(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("NoKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoValue(t *testing.T) {
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
				f: util.PassAll[string],
			},
			want: true,
		},
		{
			name: "some",
			args: args[int, string]{
				m: map[int]string{1: "1", 2: "b"},
				f: util.IsNumber,
			},
			want: false,
		},
		{
			name: "nome",
			args: args[int, string]{
				m: map[int]string{1: "a", 2: "b"},
				f: util.IsNumber,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoValue(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("NoValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeKey(t *testing.T) {
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
				f: util.PassAll[string],
			},
			want: false,
		},
		{
			name: "some",
			args: args[string, int]{
				m: map[string]int{"1": 1, "b": 2},
				f: util.IsNumber,
			},
			want: true,
		},
		{
			name: "none",
			args: args[string, int]{
				m: map[string]int{"a": 1, "b": 2},
				f: util.IsNumber,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SomeKey(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("SomeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeValue(t *testing.T) {
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
				f: util.PassAll[string],
			},
			want: false,
		},
		{
			name: "some",
			args: args[int, string]{
				m: map[int]string{1: "1", 2: "b"},
				f: util.IsNumber,
			},
			want: true,
		},
		{
			name: "nome",
			args: args[int, string]{
				m: map[int]string{1: "a", 2: "b"},
				f: util.IsNumber,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SomeValue(tt.args.m, tt.args.f); got != tt.want {
				t.Errorf("SomeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
