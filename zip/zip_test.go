package zip

import (
	"github.com/gabrielseibel1/fungo/conv"
	"reflect"
	"testing"
)

func TestSlicesToPairs(t *testing.T) {
	type args[T any, U any] struct {
		t []T
		u []U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []conv.Pair[T, U]
	}
	tests := []testCase[string, int]{
		{
			name: "both empty",
			args: args[string, int]{},
			want: []conv.Pair[string, int]{},
		},
		{
			name: "first empty",
			args: args[string, int]{
				t: nil,
				u: []int{1},
			},
			want: []conv.Pair[string, int]{},
		},
		{
			name: "second empty",
			args: args[string, int]{
				t: []string{"1"},
				u: nil,
			},
			want: []conv.Pair[string, int]{},
		},
		{
			name: "abcd123",
			args: args[string, int]{
				t: []string{"a", "b", "c", "d"},
				u: []int{1, 2, 3},
			},
			want: []conv.Pair[string, int]{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}},
		},
		{
			name: "abc1234",
			args: args[string, int]{
				t: []string{"a", "b", "c"},
				u: []int{1, 2, 3, 4},
			},
			want: []conv.Pair[string, int]{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slices(tt.args.t, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slices() = %v, want %v", got, tt.want)
			}
		})
	}
}
