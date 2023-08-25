package fungomap

import (
	"fungo/util"
	"reflect"
	"testing"
)

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
	tests := []testCase[string, util.Data]{
		{
			name: "empty/nil",
			args: args[string, util.Data]{
				m: nil,
				f: nil,
			},
			want: map[string]util.Data{},
		},
		{
			name: "remove data 1",
			args: args[string, util.Data]{
				m: map[string]util.Data{"data1": util.Data1, "data2": util.Data2},
				f: func(s string) bool {
					if s == "data1" {
						return false
					}
					return true
				},
			},
			want: map[string]util.Data{"data2": util.Data2},
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
	tests := []testCase[string, util.Data]{
		{
			name: "empty/nil",
			args: args[string, util.Data]{
				m: nil,
				f: nil,
			},
			want: map[string]util.Data{},
		},
		{
			name: "remove data 1",
			args: args[string, util.Data]{
				m: map[string]util.Data{"data1": util.Data1, "data2": util.Data2},
				f: func(data util.Data) bool {
					if data.V == util.Data1.V {
						return false
					}
					return true
				},
			},
			want: map[string]util.Data{"data2": util.Data2},
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
