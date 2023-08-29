package unzip

import (
	"github.com/gabrielseibel1/fungo/types"
	"reflect"
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
	}
	type testCase[T comparable, U any] struct {
		name  string
		args  args[T, U]
		want1 []T
		want2 []U
	}
	tests := []testCase[int, bool]{
		{
			name:  "empty",
			args:  args[int, bool]{},
			want1: []int{},
			want2: []bool{},
		},
		{
			name: "numbers",
			args: args[int, bool]{
				m: map[int]bool{1: true, 2: false, 3: true},
			},
			want1: []int{1, 2, 3},
			want2: []bool{true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := Map(tt.args.m)
			slices.Sort(got1)
			slices.SortFunc(got2, func(a, b bool) int {
				if a && !b {
					return -1
				} else if !a && b {
					return 1
				} else {
					return 0
				}
			})
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Map() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Map() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func TestPairs(t *testing.T) {
	type args[T any, U any] struct {
		p []types.Pair[T, U]
	}
	type testCase[T any, U any] struct {
		name  string
		args  args[T, U]
		want  []T
		want1 []U
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				p: []types.Pair[int, string]{},
			},
			want:  []int{},
			want1: []string{},
		},
		{
			name: "numbers",
			args: args[int, string]{
				p: []types.Pair[int, string]{{K: 0, V: "0"}, {K: 1, V: "1"}},
			},
			want:  []int{0, 1},
			want1: []string{"0", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pairs(tt.args.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pairs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pairs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
