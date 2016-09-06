package gomastermind

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	R byte = 'R'
	W byte = 'W'
	Y byte = 'Y'
	G byte = 'G'
	U byte = 'U'
	K byte = 'K'
)

func ctoi(c byte) int {
	switch c {
	case R:
		return 0
	case W:
		return 1
	case Y:
		return 2
	case G:
		return 3
	case U:
		return 4
	case K:
		return 5
	}
	return 0
}

func itoc(i int) byte {
	switch i {
	case 0:
		return R
	case 1:
		return W
	case 2:
		return Y
	case 3:
		return G
	case 4:
		return U
	case 5:
		return K
	}
	return 0
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	} else if b == 1 {
		return a
	} else if b%2 == 0 {
		return pow(a, b/2) * pow(a, b/2)
	} else {
		return pow(a, b/2) * pow(a, b/2) * a
	}

}

func Hash(guess []byte) int {
	acc := 0
	for i, g := range guess {
		acc += pow(6, i) * ctoi(g)
	}
	return acc
}

func Dehash(num int) []byte {
	code := make([]byte, 4)
	for i := 3; i >= 0; i-- {
		code[i] = itoc(num / pow(6, i))
		num = num - pow(6, i)*(num/pow(6, i))
	}
	return code
}

func SplitScore(score string) []int {
	a := make([]int, 2)
	for i := 0; i < len(score); i++ {
		switch score[i] {
		case 'x':
			a[0] += 1
		case 'o':
			a[1] += 1
		}
	}
	return a
}

func SplitGuess(guess string) []byte {
	a := make([]byte, 0, len(guess))

	guess = strings.ToUpper(guess)

	for i := 0; i < len(guess); i++ {
		if guess[i] == 'B' || guess[i] == '\n' {
			continue
		}
		a = append(a, byte(guess[i]))
	}
	return a
}

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func Judge(a, b []byte) ([]int, error) {
	snum := len(a)
	if len(a) != len(b) {
		return nil, errors.New("gomastermind: a and b must be the same length")
	}

	score := make([]int, 2)
	amark := make([]int, snum)
	bmark := make([]int, snum)

	for ia, elema := range a {
		if b[ia] == elema {
			score[0] += 1
			amark[ia] = 1
			bmark[ia] = 1
		}
	}

	for ia, elema := range a {
		if amark[ia] == 0 {
			for ib := 0; ib < snum; ib++ {
				if ib != ia && elema == b[ib] && bmark[ib] == 0 {
					score[1] += 1
					amark[ia] = 1
					bmark[ib] = 1
					break
				}
			}
		}
	}
	return score, nil
}

func JudgeFinder(guess_l []byte, score_l, pool []int, sN, cN int) []int {

	for i := 0; i < cN*cN*cN*cN; i++ {
		if pool[i] != 0 {
			score, err := Judge(guess_l, Dehash(i))

			if err != nil {
				fmt.Fprintf(os.Stderr, "JudgeFinder:", err)
				os.Exit(1)
			}

			if !reflect.DeepEqual(score, score_l) {
				pool[i] = 0
			}
		}
	}

	return pool
}
