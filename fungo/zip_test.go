package fungo

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	type args[T any, U any] struct {
		t []T
		u []U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []Pair[T, U]
	}
	tests := []testCase[string, int]{
		{
			name: "both empty",
			args: args[string, int]{},
			want: []Pair[string, int]{},
		},
		{
			name: "first empty",
			args: args[string, int]{
				t: nil,
				u: []int{1},
			},
			want: []Pair[string, int]{},
		},
		{
			name: "second empty",
			args: args[string, int]{
				t: []string{"1"},
				u: nil,
			},
			want: []Pair[string, int]{},
		},
		{
			name: "abcd123",
			args: args[string, int]{
				t: []string{"a", "b", "c", "d"},
				u: []int{1, 2, 3},
			},
			want: []Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}},
		},
		{
			name: "abc1234",
			args: args[string, int]{
				t: []string{"a", "b", "c"},
				u: []int{1, 2, 3, 4},
			},
			want: []Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.args.t, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}
