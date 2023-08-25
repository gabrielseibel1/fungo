package fungomap

import (
	"slices"
	"testing"
)

func TestToSliceK(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want []T
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				m: nil,
			},
			want: []int{},
		},
		{
			name: "some",
			args: args[int, string]{
				m: map[int]string{0: "0", 1: "2", 2: "3"},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "other",
			args: args[int, string]{
				m: map[int]string{-1: "-1", 0: "0", 1: "1"},
			},
			want: []int{-1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSliceK(tt.args.m)
			slices.Sort(tt.want)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("ToSliceK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSliceV(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[string, int]{
		{
			name: "empty",
			args: args[string, int]{
				m: nil,
			},
			want: []int{},
		},
		{
			name: "some",
			args: args[string, int]{
				m: map[string]int{"0": 0, "1": 1, "2": 2},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "other",
			args: args[string, int]{
				m: map[string]int{"-1": -1, "0": 0, "1": 1},
			},
			want: []int{-1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSliceV(tt.args.m)
			slices.Sort(tt.want)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("ToSliceV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSliceP(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want []Pair[T, U]
	}
	tests := []testCase[string, int]{
		{
			name: "empty",
			args: args[string, int]{
				m: nil,
			},
			want: []Pair[string, int]{},
		},
		{
			name: "some",
			args: args[string, int]{
				m: map[string]int{"0": 0, "1": 1, "2": 2},
			},
			want: []Pair[string, int]{{"0", 0}, {"1", 1}, {"2", 2}},
		},
		{
			name: "other",
			args: args[string, int]{
				m: map[string]int{"-1": 1, "0": 0, "1": 1},
			},
			want: []Pair[string, int]{{"-1", 1}, {"0", 0}, {"1", 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSliceP(tt.args.m)
			f := func(a, b Pair[string, int]) int {
				if a.K < b.K {
					return -1
				}
				if a.K > b.K {
					return 1
				}
				return 0
			}
			slices.SortFunc(tt.want, f)
			slices.SortFunc(got, f)
			if !slices.Equal(got, tt.want) {
				t.Errorf("ToSliceP() = %v, want %v", got, tt.want)
			}
		})
	}
}
