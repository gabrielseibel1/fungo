package fold

import (
	"math"
	"reflect"
	"testing"
)

func TestFold(t *testing.T) {
	type args[T any, U any] struct {
		s []T
		i U
		f func(U, T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, float64]{
		{
			name: "empty",
			args: args[int, float64]{
				s: nil,
				i: 0,
				f: nil,
			},
			want: 0,
		},
		{
			name: "sum",
			args: args[int, float64]{
				s: []int{-1, 0, 1, 2},
				i: -3,
				f: func(f float64, i int) float64 {
					return f + float64(i)
				},
			},
			want: -1,
		},
		{
			name: "division (no zero)",
			args: args[int, float64]{
				s: []int{5, 4, 2, 1},
				i: 10,
				f: func(f float64, i int) float64 {
					return f / float64(i)
				},
			},
			want: 0.25,
		},
		{
			name: "division (w/ zero)",
			args: args[int, float64]{
				s: []int{5, 4, 2, 1, 0},
				i: 10,
				f: func(f float64, i int) float64 {
					if i == 0 {
						return math.Inf(int(f))
					}
					return f / float64(i)
				},
			},
			want: math.Inf(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.args.s, tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
