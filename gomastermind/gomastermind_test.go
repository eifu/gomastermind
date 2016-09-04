package gomastermind

import (
	"bytes"
	"testing"
)

func TestHash(t *testing.T) {

	test1 := []byte{'R', 'W', 'Y', 'G'}

	if !bytes.Equal(test1, Dehash(Hash(test1))) {
		t.Error("Expected [82 87 89 71], got", Dehash(Hash(test1)))
	}

	for i := 0; i < 6*6*6*6; i++ {
		if Hash(Dehash(i)) != i {
			t.Error("Expected", i, ", got", Hash(Dehash(i)))
		}
	}

}

func TestElimColor(t *testing.T) {
	var colors []byte = []byte{'W', 'Y', 'G', 'U', 'K'}

	var c1 []byte = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}

	c2 := elimColor('R', c1)

	if !bytes.Equal(c2, colors) {
		t.Error("Expected", colors, ", got", c2)
	}
}
