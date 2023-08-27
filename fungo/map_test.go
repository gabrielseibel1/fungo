package fungo

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMapped(t *testing.T) {
	type args[T any, U any] struct {
		s []T
		f func(T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[Data, Record]{
		{
			name: "empty",
			args: args[Data, Record]{
				s: []Data{},
				f: DataToRecord,
			},
			want: []Record{},
		},
		{
			name: "nil",
			args: args[Data, Record]{
				s: nil,
				f: DataToRecord,
			},
			want: []Record{},
		},
		{
			name: "one",
			args: args[Data, Record]{
				s: []Data{Data1},
				f: DataToRecord,
			},
			want: []Record{Record1},
		},
		{
			name: "two",
			args: args[Data, Record]{
				s: []Data{Data1, Data2},
				f: DataToRecord,
			},
			want: []Record{Record1, Record2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MappedS(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MappedS() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	tests := []testCase[int, Data, string]{
		{
			name: "empty",
			args: args[int, Data, string]{
				m: nil,
				f: nil,
			},
			want: map[string]Data{},
		},
		{
			name: "itoa",
			args: args[int, Data, string]{
				m: map[int]Data{1: Data1, 2: Data2},
				f: strconv.Itoa,
			},
			want: map[string]Data{"1": Data1, "2": Data2},
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

func TestMappedC(t *testing.T) {
	putX := func(c chan<- int, x int) {
		for i := 0; i < x; i++ {
			c <- i
		}
		close(c)
	}
	t.Run("empty", func(t *testing.T) {
		o := make(chan int)
		m := MappedC(o, func(i int) bool { return i%2 == 0 })
		go func() { putX(o, 0) }()
		for e := range m {
			t.Errorf("expected no elements, got element %t", e)
		}
	})
	t.Run("map to is even", func(t *testing.T) {
		o := make(chan int)
		m := MappedC(o, func(i int) bool { return i%2 == 0 })
		go func() { putX(o, 10) }()
		even := true
		for e := range m {
			if e != even {
				t.Errorf("expected filtered element to be %t, got %t", even, e)
			}
			even = !even
		}
	})
}
