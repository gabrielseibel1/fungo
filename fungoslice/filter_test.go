package fungoslice

import (
	"fungo/util"
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
				f: util.PassAllData,
			},
			want: []util.Data{},
		},
		{
			name: "one element (1), pass all",
			args: args[util.Data]{
				s: []util.Data{util.Data1},
				f: util.PassAllData,
			},
			want: []util.Data{util.Data1},
		},
		{
			name: "one element (2), pass all",
			args: args[util.Data]{
				s: []util.Data{util.Data1},
				f: util.PassAllData,
			},
			want: []util.Data{util.Data1},
		},
		{
			name: "two elements, pass all",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassAllData,
			},
			want: []util.Data{util.Data1, util.Data2},
		},
		{
			name: "two elements, pass none",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassNoData,
			},
			want: []util.Data{},
		},
		{
			name: "two elements, pass first",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: func(d util.Data) bool {
					if d.V == util.Data1.V {
						return true
					} else {
						return false
					}
				},
			},
			want: []util.Data{util.Data1},
		},
		{
			name: "two elements, pass second",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: func(d util.Data) bool {
					if d.V == util.Data2.V {
						return true
					} else {
						return false
					}
				},
			},
			want: []util.Data{util.Data2},
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
