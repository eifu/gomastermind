package gomastermind

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	test1 := []byte{R, W, Y, G}

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
	var colors []byte = []byte{W, Y, G, U, K}

	var c1 []byte = []byte{R, W, Y, G, U, K}

	c2 := elimColor(R, c1)

	if !bytes.Equal(c2, colors) {
		t.Error("Expected", colors, ", got", c2)
	}
}

func TestPermute(t *testing.T) {
	a := []byte{'a', 'b', 'c', 'd'}
	var n int = 4
	permute(a, 0, n-1)

}

func TestCase1(t *testing.T) {
	fmt.Println("case1:")
	expected := []byte{R, U, Y, Y}
	pool := make([]int, 6*6*6*6)
	for i := 0; i < 6*6*6*6; i++ {
		pool[i] = 1
	}

	var guess_l []byte
	var score_l []int

	guess_l = []byte{U, W, G, Y}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{Y, G, G, Y}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{K, G, G, W}
	score_l = []int{0, 0}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{U, U, R, R}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	for i := 0; i < 6*6*6*6; i++ {
		if pool[i] != 0 {
			fmt.Println(string(Dehash(i)))
			if Hash(expected) != i {
				t.Error("Expected", expected, ", got", string(Dehash(i)))
			}
		}
	}
}

func TestCase2(t *testing.T) {
	fmt.Println("case2:")
	expected1 := []byte{G, K, G, Y}
	expected2 := []byte{G, W, G, U}
	pool := make([]int, 6*6*6*6)
	for i := 0; i < 6*6*6*6; i++ {
		pool[i] = 1
	}

	var guess_l []byte
	var score_l []int

	guess_l = []byte{R, G, Y, W}
	score_l = []int{0, 2}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{Y, K, W, U}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{G, R, R, G}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{K, U, G, R}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{W, R, K, R}
	score_l = []int{0, 1}
	pool = Finder(guess_l, score_l, pool)

	for i := 0; i < 6*6*6*6; i++ {
		if pool[i] != 0 {
			fmt.Println(string(Dehash(i)))
			if Hash(expected1) != i && Hash(expected2) != i {
				t.Error("Expected", expected1, "or", expected2, ", got", string(Dehash(i)))
			}
		}
	}
}

func TestCase3(t *testing.T) {
	fmt.Println("case3:")
	pool := make([]int, 6*6*6*6)
	for i := 0; i < 6*6*6*6; i++ {
		pool[i] = 1
	}

	var guess_l []byte
	var score_l []int

	guess_l = []byte{W, G, U, R}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{K, Y, R, U}
	score_l = []int{1, 1}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{R, R, G, Y}
	score_l = []int{0, 3}
	pool = Finder(guess_l, score_l, pool)

	guess_l = []byte{Y, R, R, W}
	score_l = []int{2, 0}
	pool = Finder(guess_l, score_l, pool)

	for i := 0; i < 6*6*6*6; i++ {
		if pool[i] != 0 {
			fmt.Println(string(Dehash(i)))

		}
	}
}
