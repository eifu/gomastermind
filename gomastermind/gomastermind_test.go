package gomastermind

import (
	"bytes"
	"testing"
)

func TestGomastermind(t *testing.T) {

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
