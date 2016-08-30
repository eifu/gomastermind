package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	R      byte         = 'R'
	W      byte         = 'W'
	Y      byte         = 'Y'
	G      byte         = 'G'
	B      byte         = 'B'
	u      byte         = 'u'
	k      byte         = 'k'
	colors []byte       = []byte{R, W, Y, G, u, k}
	ctoi   map[byte]int = map[byte]int{
		R: 0,
		W: 1,
		Y: 2,
		G: 3,
		u: 4,
		k: 5,
	}
	b1 [4][]int = [4][]int{[3]int{2, 3, 4}, [3]int{1, 3, 4}, [3]int{1, 2, 4}, [3]int{1, 2, 3}}
	b2 [6][]int = [6][]int{[2]int{1, 2}, [2]int{1, 3}, [2]int{1, 4}, [2]int{2, 3}, [2]int{2, 4}, [2]int{3, 4}}
	b3 [4][]int = [4][]int{[1]int{1}, [1]int{2}, [1]int{3}, [1]int{4}}
)

func hash(guess []byte) int {
	acc := 0
	for _, g := range guess {
		switch g {
		case R:
			acc += 0
		case W:
			acc += 1
		case Y:
			acc += 2
		case G:
			acc += 3
		case u:
			acc += 4
		case k:
			acc += 5
		}
		acc *= 6
	}
	return acc
}

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

	var pool [6 * 6 * 6 * 6]int

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

func finder(guess_l []byte, score_l, pool []int) []int {
	var index int
	if score_l[0] == 1 {
		for ib, e := range b1 {
			if score_l[0] == 0 {
				// w=0
				for _, c1 := range colors {
					for _, c2 := range colors {
						for _, c3 := range colors {
							index = math.Pow(6, ib)*ctoi[guess_l[ib]] + math.Pow(6, e[0])*ctoi[c1] + math.Pow(6, e[1])*ctoi[c2] + math.Pow(6, e[2])*ctoi[c3]
							pool[index] += 1
						}
					}
				}

			} else if score_l[1] == 1 {
				// w=1

				c1 := guess_l[e[0]]
				for _, c2 := range colors {
					for _, c3 := range colors {
						index = math.Pow(6, ib)*ctoi[guess_l[ib]] + math.Pow(6, e[0])*ctoi[c1] + math.Pow(6, e[1])*ctoi[c2] + math.Pow(6, e[2])*ctoi[c3]
						pool[index] += 1
					}
				}

				c2 := guess_l[e[1]]
				for _, c1 := range colors {
					for _, c3 := range colors {
						index = math.Pow(6, ib)*ctoi[guess_l[ib]] + math.Pow(6, e[0])*ctoi[c1] + math.Pow(6, e[1])*ctoi[c2] + math.Pow(6, e[2])*ctoi[c3]
						pool[index] += 1
					}
				}

				c3 := guess_l[e[2]]
				for _, c1 := range colors {
					for _, c2 := range colors {
						index = math.Pow(6, ib)*ctoi[guess_l[ib]] + math.Pow(6, e[0])*ctoi[c1] + math.Pow(6, e[1])*ctoi[c2] + math.Pow(6, e[2])*ctoi[c3]
						pool[index] += 1
					}
				}

			} else if score_l[1] == 2 {
				// w=2 12,13,23

			} else {
				// w=3 123
			}
		}
	} else if score_l[0] == 2 {

	} else if score_l[0] == 3 {

	} else {

	}

	return pool
}
