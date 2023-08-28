package filter

import (
	"github.com/gabrielseibel1/fungo/util"
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
	tests := []testCase[util.Data]{
		{
			name: "empty",
			args: args[util.Data]{
				f: util.PassAll[util.Data],
			},
			want: []util.Data{},
		},
		{
			name: "one element (1), pass all",
			args: args[util.Data]{
				s: []util.Data{util.Data1},
				f: util.PassAll[util.Data],
			},
			want: []util.Data{util.Data1},
		},
		{
			name: "one element (2), pass all",
			args: args[util.Data]{
				s: []util.Data{util.Data1},
				f: util.PassAll[util.Data],
			},
			want: []util.Data{util.Data1},
		},
		{
			name: "two elements, pass all",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassAll[util.Data],
			},
			want: []util.Data{util.Data1, util.Data2},
		},
		{
			name: "two elements, pass none",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassNo[util.Data],
			},
			want: []util.Data{},
		},
		{
			name: "two elements, pass first",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.Is(util.Data1),
			},
			want: []util.Data{util.Data1},
		},
		{
			name: "two elements, pass second",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.Is(util.Data2),
			},
			want: []util.Data{util.Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
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
					return s != "data1"
				},
			},
			want: map[string]util.Data{"data2": util.Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapByKeys(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapByValues() = %v, want %v", got, tt.want)
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
					return data.V != util.Data1.V
				},
			},
			want: map[string]util.Data{"data2": util.Data2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapByValues(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapByValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilteredC(t *testing.T) {
	putX := func(c chan<- int, x int) {
		for i := 0; i < x; i++ {
			c <- i
		}
		close(c)
	}
	t.Run("empty", func(t *testing.T) {
		unfiltered := make(chan int)
		filtered := Channel(unfiltered, func(i int) bool { return i%2 == 0 })
		go func() { putX(unfiltered, 0) }()
		for e := range filtered {
			t.Errorf("expected no elements, got element %d", e)
		}
	})
	t.Run("filter even", func(t *testing.T) {
		unfiltered := make(chan int)
		filtered := Channel(unfiltered, func(i int) bool { return i%2 == 0 })
		go func() { putX(unfiltered, 4) }()
		i := 0
		for e := range filtered {
			if e != i {
				t.Errorf("expected filtered element to be %d, got %d", i, e)
			}
			i += 2
		}
	})
	t.Run("filter even then div by 4", func(t *testing.T) {
		unfiltered := make(chan int)
		filtered1 := Channel(unfiltered, func(i int) bool { return i%2 == 0 })
		filtered2 := Channel(filtered1, func(i int) bool { return i%4 == 0 })
		go func() { putX(unfiltered, 21) }()
		i := 0
		for e := range filtered2 {
			if e != i {
				t.Errorf("expected filtered element to be %d, got %d", i, e)
			}
			i += 4
		}
	})
	// TODO test buffered
}
