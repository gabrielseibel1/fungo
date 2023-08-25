package fungomap

import (
	"fungo/util"
	"reflect"
	"strconv"
	"testing"
)

func TestMappedK(t *testing.T) {
	type args[T comparable, V any, U comparable] struct {
		m map[T]V
		f func(T) U
	}
	type testCase[T comparable, V any, U comparable] struct {
		name string
		args args[T, V, U]
		want map[U]V
	}
	tests := []testCase[int, util.Data, string]{
		{
			name: "empty",
			args: args[int, util.Data, string]{
				m: nil,
				f: nil,
			},
			want: map[string]util.Data{},
		},
		{
			name: "itoa",
			args: args[int, util.Data, string]{
				m: map[int]util.Data{1: util.Data1, 2: util.Data2},
				f: strconv.Itoa,
			},
			want: map[string]util.Data{"1": util.Data1, "2": util.Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MappedK(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MappedK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMappedV(t *testing.T) {
	type args[T comparable, U any, V any] struct {
		m map[T]U
		f func(U) V
	}
	type testCase[T comparable, U any, V any] struct {
		name string
		args args[T, U, V]
		want map[T]V
	}
	tests := []testCase[string, float32, int32]{
		{
			name: "empty",
			args: args[string, float32, int32]{
				m: nil,
				f: nil,
			},
			want: map[string]int32{},
		},
		{
			name: "truncate",
			args: args[string, float32, int32]{
				m: map[string]float32{"1.1": 1.1, "2.2": 2.2},
				f: func(f float32) int32 { return int32(f) },
			},
			want: map[string]int32{"1.1": 1, "2.2": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MappedV(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MappedV() = %v, want %v", got, tt.want)
			}
		})
	}
}
