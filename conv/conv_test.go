package conv

import (
	"reflect"
	"slices"
	"strconv"
	"testing"
)

func TestSliceToMapWithIndices(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want map[int]T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{
				s: nil,
			},
			want: map[int]string{},
		},
		{
			name: "some elements",
			args: args[string]{
				s: []string{"0", "1", "2"},
			},
			want: map[int]string{0: "0", 1: "1", 2: "2"},
		},
		{
			name: "other elements",
			args: args[string]{
				s: []string{"-1", "0", "1"},
			},
			want: map[int]string{0: "-1", 1: "0", 2: "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToMapWithIndices(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMapWithIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToMapKeys(t *testing.T) {
	type args[T comparable, U any] struct {
		s []T
		v func(T) U
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want map[T]U
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				s: nil,
				v: strconv.Itoa,
			},
			want: map[int]string{},
		},
		{
			name: "some elements",
			args: args[int, string]{
				s: []int{0, 1, 2},
				v: strconv.Itoa,
			},
			want: map[int]string{0: "0", 1: "1", 2: "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToMapKeys(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToMapValues(t *testing.T) {
	type args[T any, U comparable] struct {
		s []T
		k func(T) U
	}
	type testCase[T any, U comparable] struct {
		name string
		args args[T, U]
		want map[U]T
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				s: nil,
				k: strconv.Itoa,
			},
			want: map[string]int{},
		},
		{
			name: "some elements",
			args: args[int, string]{
				s: []int{0, 1, 2},
				k: strconv.Itoa,
			},
			want: map[string]int{"0": 0, "1": 1, "2": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToMapValues(tt.args.s, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMapValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysToSlice(t *testing.T) {
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
			got := MapKeysToSlice(tt.args.m)
			slices.Sort(tt.want)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MapKeysToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapValuesToSlice(t *testing.T) {
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
			got := MapValuesToSlice(tt.args.m)
			slices.Sort(tt.want)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MapValuesToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToPairs(t *testing.T) {
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
			got := MapToPairs(tt.args.m)
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
				t.Errorf("MapToPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
