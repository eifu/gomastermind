package main

import (
	"./gomastermind"
	"bufio"
	"fmt"
	"os"
)

const (
	R byte = 'R'
	W byte = 'W'
	Y byte = 'Y'
	G byte = 'G'
	U byte = 'U'
	K byte = 'K'
)

func main() {
	fmt.Println("welcome to master mind game.")
	reader := bufio.NewReader(os.Stdin)
	var guesscount int

	var guess_l []byte
	var score_l []int

	pool := make([]int, 6*6*6*6)
	for i := 0; i < 6*6*6*6; i++ {
		pool[i] = 1
	}
	for {
		guesscount += 1
		fmt.Printf("Guess %d\n", guesscount)
		fmt.Println("enter color: red(R), white(W), yellow(Y), green(G), blue(Bu), and black(Bk).")
		guess, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader: %v\n", guess, err)
			os.Exit(1)
		}
		guess_l = gomastermind.SplitGuess(guess)
		fmt.Println(guess_l)
		fmt.Println("enter score: x for right color, right position, o for right color but in the wrong position.")
		score, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader: %v\n", score, err)
			os.Exit(1)
		}

		score_l = gomastermind.SplitScore(score)
		fmt.Println(score_l)

		pool = gomastermind.Finder(guess_l, score_l, pool)
		fmt.Println(pool)
		var count int
		for index, e := range pool {
			if e != 0 {
				count += 1
				fmt.Println(index, string(gomastermind.Dehash(index)))
			}
		}
		fmt.Println(count, "cases found")

	}
}
