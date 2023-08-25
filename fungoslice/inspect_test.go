package fungoslice

import (
	"fungo/util"
	"testing"
)

func TestSome(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[util.Data]{
		{
			name: "empty/nil",
			args: args[util.Data]{
				s: nil,
				f: util.PassAll[util.Data],
			},
			want: false,
		},
		{
			name: "some",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.Is(util.Data1),
			},
			want: true,
		},
		{
			name: "none",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassNo[util.Data],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[util.Data]{
		{
			name: "empty/nil",
			args: args[util.Data]{
				s: nil,
				f: util.PassAll[util.Data],
			},
			want: true,
		},
		{
			name: "some",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.Is(util.Data1),
			},
			want: false,
		},
		{
			name: "none",
			args: args[util.Data]{
				s: []util.Data{util.Data1, util.Data2},
				f: util.PassNo[util.Data],
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := None(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}
