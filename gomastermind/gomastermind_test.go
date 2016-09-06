package gomastermind

import (
	"bytes"
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

func TestJudgeFinderCase1(t *testing.T) {

	expected := []byte{R, U, Y, Y}
	pool := make([]int, 6*6*6*6)
	for i := 0; i < 6*6*6*6; i++ {
		pool[i] = 1
	}

	var guess_l []byte
	var score_l []int

	guess_l = []byte{U, W, G, Y}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{Y, G, G, Y}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{K, G, G, W}
	score_l = []int{0, 0}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{U, U, R, R}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	for i := 0; i < 6*6*6*6; i++ {
		if pool[i] != 0 {
			if Hash(expected) != i {
				t.Error("Expected", expected, ", got", string(Dehash(i)))
			}
		}
	}
}

func TestJudgeFinderCase2(t *testing.T) {

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
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{Y, K, W, U}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{G, R, R, G}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{K, U, G, R}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{W, R, K, R}
	score_l = []int{0, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	for i := 0; i < 6*6*6*6; i++ {
		if pool[i] != 0 {
			if Hash(expected1) != i && Hash(expected2) != i {
				t.Error("Expected", expected1, "or", expected2, ", got", string(Dehash(i)))
			}
		}
	}
}

func TestJudgeFinderCase3(t *testing.T) {
	expected := []byte{Y, G, R, G}
	pool := make([]int, 6*6*6*6)
	for i := 0; i < 6*6*6*6; i++ {
		pool[i] = 1
	}

	var guess_l []byte
	var score_l []int

	guess_l = []byte{W, G, U, R}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{K, Y, R, U}
	score_l = []int{1, 1}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{R, R, G, Y}
	score_l = []int{0, 3}
	pool = JudgeFinder(guess_l, score_l, pool)

	guess_l = []byte{Y, R, R, W}
	score_l = []int{2, 0}
	pool = JudgeFinder(guess_l, score_l, pool)

	for i := 0; i < 6*6*6*6; i++ {
		if pool[i] != 0 {
			if Hash(expected) != i {
				t.Error("Expected", expected, ", got", string(Dehash(i)))
			}
		}
	}
}

func TestJudge(t *testing.T) {
	var a, b []byte
	var j []int

	a = []byte{R, W, Y, G}
	b = []byte{W, R, Y, G}
	j = Judge(a, b)

	if j[0] != 2 || j[1] != 2 {
		t.Error(string(a), string(b), "expected [2, 2] but", j)
	}

	a = []byte{W, W, Y, G}
	b = []byte{W, R, Y, G}
	j = Judge(a, b)

	if j[0] != 3 || j[1] != 0 {
		t.Error(string(a), string(b), "expected [3, 0] but", j)
	}

	a = []byte{G, G, G, G}
	b = []byte{W, R, Y, G}
	j = Judge(a, b)

	if j[0] != 1 || j[1] != 0 {
		t.Error(string(a), string(b), "expected [1, 0] but", j)
	}

	a = []byte{U, U, U, G}
	b = []byte{W, R, Y, G}
	j = Judge(a, b)

	if j[0] != 1 || j[1] != 0 {
		t.Error(string(a), string(b), "expected [1, 0] but", j)
	}

	a = []byte{R, Y, Y, R}
	b = []byte{W, R, Y, G}
	j = Judge(a, b)

	if j[0] != 1 || j[1] != 1 {
		t.Error(string(a), string(b), "expected [1, 1] but", j)
	}

	a = []byte{R, U, Y, Y}
	b = []byte{U, W, G, Y}
	j = Judge(a, b)

	if j[0] != 1 || j[1] != 1 {
		t.Error(string(a), string(b), "expected [1, 1] but", j)
	}

	a = []byte{R, U, Y, Y}
	b = []byte{Y, G, G, Y}
	j = Judge(a, b)

	if j[0] != 1 || j[1] != 1 {
		t.Error(string(a), string(b), "expected [1, 1] but", j)
	}

	a = []byte{R, U, Y, Y}
	b = []byte{U, U, R, R}
	j = Judge(a, b)

	if j[0] != 1 || j[1] != 1 {
		t.Error(string(a), string(b), "expected [1, 1] but", j)
	}

	a = []byte{K, K, K, R}
	b = []byte{W, R, Y, G}
	j = Judge(a, b)

	if j[0] != 0 || j[1] != 1 {
		t.Error(string(a), string(b), "expected [0, 1] but", j)
	}
}
