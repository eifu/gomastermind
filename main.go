package main

import (
	"./gomastermind"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("welcome to master mind game.")
	reader := bufio.NewReader(os.Stdin)
	var guesscount, sN, cN int

	var guess_l []byte
	var score_l []int

	fmt.Println("enter the num of sequence")
	sNbyte, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "reader:1 %v\n", sNbyte, err)
		os.Exit(1)
	}

	fmt.Println("enter the num of color")
	cNbyte, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "reader:3 %v\n", cNbyte, err)
		os.Exit(1)
	}

	sN, err = strconv.Atoi(string(sNbyte[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "reader:5 %v\n", sNbyte, err)
		os.Exit(1)
	}
	cN, err = strconv.Atoi(string(cNbyte[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "reader:6 %v\n", cNbyte, err)
		os.Exit(1)
	}

	pool := make([]int, cN*cN*cN*cN)
	for i := 0; i < cN*cN*cN*cN; i++ {
		pool[i] = 1
	}

	for {
		guesscount += 1
		fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-")
		fmt.Printf("Guess %d\n", guesscount)
		fmt.Println("enter color: red(R), white(W), yellow(Y), green(G), blue(Bu), and black(Bk).")
		guess, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader: %v\n", guess, err)
			os.Exit(1)
		}
		guess_l = gomastermind.SplitGuess(guess)

		fmt.Println("enter score: x for right color, right position, o for right color but in the wrong position.")
		score, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "reader: %v\n", score, err)
			os.Exit(1)
		}

		score_l = gomastermind.SplitScore(score)

		pool = gomastermind.JudgeFinder(guess_l, score_l, pool, sN, cN)

		count := 0
		for index, e := range pool {
			if e != 0 {
				count += 1
				fmt.Println(index, string(gomastermind.Dehash(index)))
			}
		}

		fmt.Println(count, "cases found")

	}
}
