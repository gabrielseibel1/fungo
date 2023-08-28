package main

import (
	"github.com/gabrielseibel1/fungo/apply"
	"github.com/gabrielseibel1/fungo/check"
	"github.com/gabrielseibel1/fungo/conv"
	"github.com/gabrielseibel1/fungo/filter"
	"github.com/gabrielseibel1/fungo/fold"
	"github.com/gabrielseibel1/fungo/zip"
	"log/slog"
	"slices"
	"strings"
)

func main() {
	d := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4}
	m1 := conv.SliceToMapWithIndices(d)
	m2 := apply.ToValues(m1, numText)
	s1 := conv.MapKeysToSlice(m2)
	slices.Sort(s1)
	s2 := conv.MapValuesToSlice(m2)
	slices.Sort(s2)
	z := zip.Slices(s1, s2)
	a := apply.ToSlice(z, repeatString)
	containsO := func(s string) bool { return strings.Contains(s, "o") }
	f := filter.Slice(a, containsO)
	b1 := check.Some(f, containsO)
	b2 := fold.Slice(f, b1, func(b bool, s string) bool {
		return b && strings.Contains(s, "o")
	})
	bs := conv.SliceToMapKeys([]bool{b1, b2}, func(t bool) int {
		if t {
			return 1
		}
		return 0
	})
	at := check.NoValue(bs, func(i int) bool {
		return i == 0
	})
	if !at {
		slog.Error("failure",
			"d", d,
			"m1", m1,
			"m2", m2,
			"s1", s1,
			"s2", s2,
			"z", z,
			"a", a,
			"f", f,
			"b1", b1,
			"b2", b2,
			"bs", bs,
			"at", at,
		)
		panic("failure")
	}
	println("success")
}

func numText(i int) string {
	switch i {
	case -4:
		return "mfour"
	case -3:
		return "mthree"
	case -2:
		return "mtwo"
	case -1:
		return "mone"
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return "?"
	}
}
func repeatString(t conv.Pair[int, string]) string {
	s := t.V
	for i := 0; i < t.K; i++ {
		s += s
	}
	return s
}
