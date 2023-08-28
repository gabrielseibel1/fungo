package util

import (
	"reflect"
	"testing"
)

func TestDataToRecord(t *testing.T) {
	type args struct {
		d Data
	}
	tests := []struct {
		name string
		args args
		want Record
	}{
		{
			name: "data 1 to record 1",
			args: args{
				d: Data1,
			},
			want: Record1,
		},
		{
			name: "data 2 to record 2",
			args: args{
				d: Data2,
			},
			want: Record2,
		},
		{
			name: "other data to zero record",
			args: args{
				d: Data{V: "other"},
			},
			want: Record{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DataToRecord(tt.args.d); !reflect.DeepEqual(got, tt.want) {
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
			if got := IsNumber(tt.args.s); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
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
			if got := PassAll(tt.args.in0); got != tt.want {
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
			if got := PassNo(tt.args.in0); got != tt.want {
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
	tests := []testCase[Data]{
		{
			name: "data 1",
			args: args[Data]{
				t: Data1,
			},
			u:    Data1,
			want: true,
		},
		{
			name: "data 2",
			args: args[Data]{
				t: Data2,
			},
			u:    Data2,
			want: true,
		},
		{
			name: "mismatch",
			args: args[Data]{
				t: Data1,
			},
			u:    Data2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.t)(tt.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
