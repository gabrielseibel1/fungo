package fungoslice

import (
	"fungo/util"
	"reflect"
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
			if got := Mapped(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mapped() = %v, want %v", got, tt.want)
			}
		})
	}
}
