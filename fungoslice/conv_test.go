package fungoslice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestToMapI(t *testing.T) {
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
			if got := ToMapI(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapK(t *testing.T) {
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
			if got := ToMapK(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapV(t *testing.T) {
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
			if got := ToMapV(tt.args.s, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapV() = %v, want %v", got, tt.want)
			}
		})
	}
}
