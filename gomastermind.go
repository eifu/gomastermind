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
	case u:
		return 4
	case k:
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
		return u
	case 5:
		return k
	}
	return 0
}

func hash(guess []byte) int {
	acc := 0

	for _, g := range guess {
		acc += ctoi(g)
		acc *= 6
	}
	return acc
}

func dehash(num int) []byte {
	code := make([]byte, 4)
	for i := 3; i >= 0; i-- {
		code[i] = itoc(num / pow(6, i))
		num = num - pow(6, i)*(num/pow(6, i))
	}
	return code

}

func splitScore(score string) []int {
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

		pool = finder(guess_l, score_l, pool)
		fmt.Println(pool)

		for index, e := range pool {
			if e != 0 {
				fmt.Println(index, string(dehash(index)))
			}
		}

	}
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

func finder(guess_l []byte, score_l, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors [6]byte = [6]byte{'R', 'W', 'Y', 'G', 'u', 'k'}
	var ctoi map[byte]int = map[byte]int{
		R: 0,
		W: 1,
		Y: 2,
		G: 3,
		u: 4,
		k: 5,
	}
	var b1 [4][]int = [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	//	var b2 [6][]int = [6][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}}
	//	var b3 [4][]int = [4][]int{[]int{1}, []int{2}, []int{3}, []int{4}}
	var index int
	if score_l[0] == 1 {
		// b=1
		for ib, e := range b1 {
			if score_l[1] == 0 {
				// w=0
				for _, c1 := range colors { // 1st
					for _, c2 := range colors { // 2nd
						for _, c3 := range colors { // 3rd
							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[0]) * ctoi[c1]
							index += pow(6, e[1]) * ctoi[c2]
							index += pow(6, e[2]) * ctoi[c3]
							if pool[index] != 0 {
								newpool[index] = 1
							}
						}
					}
				}

			} else if score_l[1] == 1 {
				// w=1

				for pos := 0; pos < 3; pos++ {
					c0 := guess_l[e[pos]]
					for _, c1 := range colors {
						for _, c2 := range colors {

							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[0]) * ctoi[c0]
							index += pow(6, e[1]) * ctoi[c1]
							index += pow(6, e[2]) * ctoi[c2]
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[0]) * ctoi[c0]
							index += pow(6, e[2]) * ctoi[c1]
							index += pow(6, e[1]) * ctoi[c2]
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[1]) * ctoi[c0]
							index += pow(6, e[0]) * ctoi[c1]
							index += pow(6, e[2]) * ctoi[c2]
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[1]) * ctoi[c0]
							index += pow(6, e[2]) * ctoi[c1]
							index += pow(6, e[0]) * ctoi[c2]
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[2]) * ctoi[c0]
							index += pow(6, e[0]) * ctoi[c1]
							index += pow(6, e[1]) * ctoi[c2]
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi[guess_l[ib]]
							index += pow(6, e[2]) * ctoi[c0]
							index += pow(6, e[1]) * ctoi[c1]
							index += pow(6, e[0]) * ctoi[c2]
							if pool[index] != 0 {
								newpool[index] = 1
							}

						}
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

	return newpool
}
