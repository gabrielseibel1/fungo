package zip_test

import (
	"github.com/gabrielseibel1/fungo/types"
	"github.com/gabrielseibel1/fungo/zip"
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
		want []types.Pair[T, U]
	}
	tests := []testCase[string, int]{
		{
			name: "both empty",
			args: args[string, int]{},
			want: []types.Pair[string, int]{},
		},
		{
			name: "first empty",
			args: args[string, int]{
				t: nil,
				u: []int{1},
			},
			want: []types.Pair[string, int]{},
		},
		{
			name: "second empty",
			args: args[string, int]{
				t: []string{"1"},
				u: nil,
			},
			want: []types.Pair[string, int]{},
		},
		{
			name: "abcd123",
			args: args[string, int]{
				t: []string{"a", "b", "c", "d"},
				u: []int{1, 2, 3},
			},
			want: []types.Pair[string, int]{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}},
		},
		{
			name: "abc1234",
			args: args[string, int]{
				t: []string{"a", "b", "c"},
				u: []int{1, 2, 3, 4},
			},
			want: []types.Pair[string, int]{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zip.SlicesToPairs(tt.args.t, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlicesToPairss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlicesToMap(t *testing.T) {
	type args[T comparable, U any] struct {
		t []T
		u []U
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want map[T]U
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{},
			want: map[int]string{},
		},
		{
			name: "unique numbers",
			args: args[int, string]{
				t: []int{1, 2, 3},
				u: []string{"1", "2", "3"},
			},
			want: map[int]string{1: "1", 2: "2", 3: "3"},
		},
		{
			name: "repeated numbers",
			args: args[int, string]{
				t: []int{1, 2, 2},
				u: []string{"1", "2", "3"},
			},
			want: map[int]string{1: "1", 2: "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zip.SlicesToMap(tt.args.t, tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlicesToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
