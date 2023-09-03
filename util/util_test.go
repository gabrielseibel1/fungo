package util_test

import (
	"github.com/gabrielseibel1/fungo/util"
	"reflect"
	"testing"
)

func TestDataToRecord(t *testing.T) {
	type args struct {
		d util.Data
	}
	tests := []struct {
		name string
		args args
		want util.Record
	}{
		{
			name: "data 1 to record 1",
			args: args{
				d: util.Data1,
			},
			want: util.Record1,
		},
		{
			name: "data 2 to record 2",
			args: args{
				d: util.Data2,
			},
			want: util.Record2,
		},
		{
			name: "other data to zero record",
			args: args{
				d: util.Data{V: "other"},
			},
			want: util.Record{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.DataToRecord(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataToRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "42",
			args: args{
				s: "42",
			},
			want: true,
		},
		{
			name: "a",
			args: args{
				s: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.IsInt(tt.args.s); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassAll(t *testing.T) {
	type args[T any] struct {
		in0 T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[bool]{
		{
			name: "true",
			args: args[bool]{
				in0: true,
			},
			want: true,
		},
		{
			name: "false",
			args: args[bool]{
				in0: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.PassAll(tt.args.in0); got != tt.want {
				t.Errorf("PassAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassNo(t *testing.T) {
	type args[T any] struct {
		in0 T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[bool]{
		{
			name: "true",
			args: args[bool]{
				in0: true,
			},
			want: false,
		},
		{
			name: "false",
			args: args[bool]{
				in0: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.PassNo(tt.args.in0); got != tt.want {
				t.Errorf("PassNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs(t *testing.T) {
	type args[T any] struct {
		t T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		u    T
		want bool
	}
	tests := []testCase[util.Data]{
		{
			name: "data 1",
			args: args[util.Data]{
				t: util.Data1,
			},
			u:    util.Data1,
			want: true,
		},
		{
			name: "data 2",
			args: args[util.Data]{
				t: util.Data2,
			},
			u:    util.Data2,
			want: true,
		},
		{
			name: "mismatch",
			args: args[util.Data]{
				t: util.Data1,
			},
			u:    util.Data2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.Is(tt.args.t)(tt.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
