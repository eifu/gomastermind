package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	R byte = 'R'
	W byte = 'W'
	Y byte = 'Y'
	G byte = 'G'
	B byte = 'B'
	u byte = 'u'
	k byte = 'k'
)

func splitScore(score string) []int {
	a := make([]int, 2)
	for i := 0; i < len(score); i++ {
		switch score[i] {
		case 'x':
			a[0] += 1
		default:
			a[1] += 1
		}
	}
	return a
}

func splitGuess(guess string) []byte {
	a := make([]byte, 0, len(guess))
	fmt.Println(strings.Split(guess, ""))
	for i := 0; i < len(guess); i++ {
		fmt.Println(a, i, len(guess))
		switch guess[i] {
		case R:
			a = append(a, R)
		case W:
			a = append(a, W)
		case Y:
			a = append(a, Y)
		case G:
			a = append(a, G)
		case B:
			if guess[i+1] == u {
				a = append(a, u)
			} else {
				a = append(a, k)
			}
			i += 1
		default: // guess includes \n
			return a
		}
	}
	return a
}

func main() {
	fmt.Println("welcome to master mind game.")
	reader := bufio.NewReader(os.Stdin)
	var guesscount int

	var guess_l []byte
	var score_l []int

	for {
		guesscount += 1
		fmt.Printf("Guess %d\n", guesscount)
		fmt.Println("enter color: red(R), white(W), yellow(Y), green(G), blue(Bu), and black(Bk).")
		guess, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader: %v\n", guess, err)
			os.Exit(1)
		}
		guess_l = splitGuess(guess)
		fmt.Println(guess_l)
		fmt.Println("enter score: x for right color, right position, o for right color but in the wrong position.")
		score, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader: %v\n", score, err)
			os.Exit(1)
		}

		score_l = splitScore(score)
		fmt.Println(score_l)

	}
}
