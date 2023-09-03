package types_test

import (
	"github.com/gabrielseibel1/fungo/types"
	"strconv"
	"testing"
)

func TestPair(t *testing.T) {
	p := types.Pair[int, string]{
		K: 42,
		V: "42",
	}
	i, err := strconv.Atoi(p.V)
	if err != nil {
		t.Error(err)
	}
	if i != p.K {
		t.Errorf("parsed unequal number, expected %d, got %d", p.K, i)
	}
}
