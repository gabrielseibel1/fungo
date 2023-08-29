package apply

import (
	"github.com/gabrielseibel1/fungo/util"
	"reflect"
	"strconv"
	"testing"
)

func TestToSlice(t *testing.T) {
	type args[T any, U any] struct {
		s []T
		f func(T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[util.Data, util.Record]{
		{
			name: "empty",
			args: args[util.Data, util.Record]{
				s: []util.Data{},
				f: util.DataToRecord,
			},
			want: []util.Record{},
		},
		{
			name: "nil",
			args: args[util.Data, util.Record]{
				s: nil,
				f: util.DataToRecord,
			},
			want: []util.Record{},
		},
		{
			name: "one",
			args: args[util.Data, util.Record]{
				s: []util.Data{util.Data1},
				f: util.DataToRecord,
			},
			want: []util.Record{util.Record1},
		},
		{
			name: "two",
			args: args[util.Data, util.Record]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.DataToRecord,
			},
			want: []util.Record{util.Record1, util.Record2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSlice(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToKeys(t *testing.T) {
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
			if got := ToKeys(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToValues(t *testing.T) {
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
			if got := ToValues(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToChannel(t *testing.T) {
	putX := func(c chan<- int, x int) {
		for i := 0; i < x; i++ {
			c <- i
		}
		close(c)
	}
	t.Run("empty", func(t *testing.T) {
		o := make(chan int)
		m := ToChannel(o, func(i int) bool { return i%2 == 0 })
		go func() { putX(o, 0) }()
		for e := range m {
			t.Errorf("expected no elements, got element %t", e)
		}
	})
	t.Run("map to is even", func(t *testing.T) {
		o := make(chan int)
		m := ToChannel(o, func(i int) bool { return i%2 == 0 })
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
