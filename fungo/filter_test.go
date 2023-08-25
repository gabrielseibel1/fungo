package fungo

import (
	"reflect"
	"testing"
)

func TestFiltered(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[Data]{
		{
			name: "empty",
			args: args[Data]{
				f: PassAll[Data],
			},
			want: []Data{},
		},
		{
			name: "one element (1), pass all",
			args: args[Data]{
				s: []Data{Data1},
				f: PassAll[Data],
			},
			want: []Data{Data1},
		},
		{
			name: "one element (2), pass all",
			args: args[Data]{
				s: []Data{Data1},
				f: PassAll[Data],
			},
			want: []Data{Data1},
		},
		{
			name: "two elements, pass all",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: PassAll[Data],
			},
			want: []Data{Data1, Data2},
		},
		{
			name: "two elements, pass none",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: PassNo[Data],
			},
			want: []Data{},
		},
		{
			name: "two elements, pass first",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: Is(Data1),
			},
			want: []Data{Data1},
		},
		{
			name: "two elements, pass second",
			args: args[Data]{
				s: []Data{Data1, Data2},
				f: Is(Data2),
			},
			want: []Data{Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filtered(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filtered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilteredByK(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
		f func(T) bool
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want map[T]U
	}
	tests := []testCase[string, Data]{
		{
			name: "empty/nil",
			args: args[string, Data]{
				m: nil,
				f: nil,
			},
			want: map[string]Data{},
		},
		{
			name: "remove data 1",
			args: args[string, Data]{
				m: map[string]Data{"data1": Data1, "data2": Data2},
				f: func(s string) bool {
					if s == "data1" {
						return false
					}
					return true
				},
			},
			want: map[string]Data{"data2": Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilteredByK(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilteredByV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilteredByV(t *testing.T) {
	type args[T comparable, U any] struct {
		m map[T]U
		f func(U) bool
	}
	type testCase[T comparable, U any] struct {
		name string
		args args[T, U]
		want map[T]U
	}
	tests := []testCase[string, Data]{
		{
			name: "empty/nil",
			args: args[string, Data]{
				m: nil,
				f: nil,
			},
			want: map[string]Data{},
		},
		{
			name: "remove data 1",
			args: args[string, Data]{
				m: map[string]Data{"data1": Data1, "data2": Data2},
				f: func(data Data) bool {
					if data.V == Data1.V {
						return false
					}
					return true
				},
			},
			want: map[string]Data{"data2": Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilteredByV(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilteredByV() = %v, want %v", got, tt.want)
			}
		})
	}
}
