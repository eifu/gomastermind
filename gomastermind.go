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

func split(s string) []byte {
	a := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = s[i]
	}
	return a
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
	fmt.Println(split(guess))

	guess = strings.ToUpper(guess)
	guessSplit := split(guess)
	for i := 0; i < len(guess); i++ {
		fmt.Println(a, i, len(guess))

		if guessSplit[i] == 'B' || guessSplit[i] == '\n' {
			continue
		}
		a = append(a, byte(guessSplit[i]))
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

func elimColor(c byte, colors []byte) []byte {
	j := 0
	for i := 0; i < len(colors); i++ {
		if colors[i] == c {
			colors = append(colors[:(i-j)], colors[(i-j+1):]...)
			j += 1
		}
	}
	return colors
}

func finder(guess_l []byte, score_l, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors []byte = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
	var c0, c1, c2, c3 byte
	var b0 [1][]int = [1][]int{[]int{0, 1, 2, 3}}
	var b1 [4][]int = [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	//	var b2 [6][]int = [6][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}}
	//var b3 [4][]int = [4][]int{[]int{1}, []int{2}, []int{3}, []int{4}}
	var index int
	if score_l[0] == 0 {
		// b=0
		for _, e := range b0 {
			if score_l[1] == 0 {
				// w=0
				for i := 0; i < 4; i++ {
					colors = elimColor(guess_l[i], colors)
				}
				for _, c0 := range colors {
					for _, c1 := range colors {
						for _, c2 := range colors {
							for _, c3 := range colors {
								index = pow(6, 0) * ctoi(c0)
								index += pow(6, 1) * ctoi(c1)
								index += pow(6, 2) * ctoi(c2)
								index += pow(6, 3) * ctoi(c3)
								if pool[index] != 0 {
									newpool[index] = 1
								}
							}
						}
					}
				}
			} else if score_l[1] == 1 {
				// w=1
				// pick a color from guess_l, then remove all colors in guess_l from colors
				for _, pos := range e {
					c0 := guess_l[pos]       // pick a color from guess_l
					for i := 0; i < 4; i++ { // remove all colors from colors
						colors = elimColor(guess_l[i], colors)
					}
					for _, c1 := range colors {
						for _, c2 := range colors {
							for _, c3 := range colors {
								for i := 1; i < 4; i++ {
									if c0 != guess_l[(pos+1)%4] {
										index = pow(6, (pos+i)%4) * ctoi(c0)
										index += pow(6, (pos+i+1)%4) * ctoi(c1)
										index += pow(6, (pos+i+2)%4) * ctoi(c2)
										index += pow(6, (pos+i+3)%4) * ctoi(c3)
										if pool[index] != 0 {
											newpool[index] = 1
										}
									}
								}
							}
						}
					}
				}

			} else if score_l[1] == 2 {
				// w=2
				// TODO: not in the right position.
				// 01, 02, 03, 12, 13, 23

				// (0, 1) fixed, (2, 3) change
				c0 = guess_l[0]
				c1 = guess_l[1]
				for _, c2 := range colors {
					for _, c3 := range colors {
						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

					}
				}

				// 01, 02, 03, 12, 13, 23
				// (0, 2) fixed, (1, 3) change
				c0 = guess_l[0]
				c2 = guess_l[2]
				for _, c1 := range colors {
					for _, c3 := range colors {
						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

					}
				}

				// 01, 02, 03, 12, 13, 23
				// (0, 3) fixed, (1, 2) change
				c0 = guess_l[0]
				c3 = guess_l[3]
				for _, c1 := range colors {
					for _, c2 := range colors {
						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

					}
				}

				// 01, 02, 03, 12, 13, 23
				// (1, 2) fixed, (0, 3) change
				c1 = guess_l[1]
				c2 = guess_l[2]
				for _, c0 := range colors {
					for _, c3 := range colors {
						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

					}
				}

				// 01, 02, 03, 12, 13, 23
				// (1, 3) fixed, (0, 2) change
				c1 = guess_l[1]
				c3 = guess_l[3]
				for _, c0 := range colors {
					for _, c2 := range colors {
						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

					}
				}

				// 01, 02, 03, 12, 13, 23
				// (2, 3) fixed, (0, 1) change
				c2 = guess_l[2]
				c3 = guess_l[3]
				for _, c0 := range colors {
					for _, c1 := range colors {
						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 0) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 2) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 1) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 3) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 2) * ctoi(c0)
						index += pow(6, 3) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 0) * ctoi(c1)
						index += pow(6, 1) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 1) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 2) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

						index = pow(6, 3) * ctoi(c0)
						index += pow(6, 2) * ctoi(c1)
						index += pow(6, 0) * ctoi(c2)
						index += pow(6, 1) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}

					}
				}

			}

		}
	} else if score_l[0] == 1 {
		// b=1
		for ib, e := range b1 {
			if score_l[1] == 0 {
				// w=0
				for _, c0 := range colors { // 1st
					for _, c1 := range colors { // 2nd
						for _, c2 := range colors { // 3rd
							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[0]) * ctoi(c0)
							index += pow(6, e[1]) * ctoi(c1)
							index += pow(6, e[2]) * ctoi(c2)
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

							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[0]) * ctoi(c0)
							index += pow(6, e[1]) * ctoi(c1)
							index += pow(6, e[2]) * ctoi(c2)
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[0]) * ctoi(c0)
							index += pow(6, e[2]) * ctoi(c1)
							index += pow(6, e[1]) * ctoi(c2)
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[1]) * ctoi(c0)
							index += pow(6, e[0]) * ctoi(c1)
							index += pow(6, e[2]) * ctoi(c2)
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[1]) * ctoi(c0)
							index += pow(6, e[2]) * ctoi(c1)
							index += pow(6, e[0]) * ctoi(c2)
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[2]) * ctoi(c0)
							index += pow(6, e[0]) * ctoi(c1)
							index += pow(6, e[1]) * ctoi(c2)
							if pool[index] != 0 {
								newpool[index] = 1
							}

							index = pow(6, ib) * ctoi(guess_l[ib])
							index += pow(6, e[2]) * ctoi(c0)
							index += pow(6, e[1]) * ctoi(c1)
							index += pow(6, e[0]) * ctoi(c2)
							if pool[index] != 0 {
								newpool[index] = 1
							}

						}
					}

				}

			} else if score_l[1] == 2 {
				// w=2 01, 02, 12
				c0 := guess_l[e[0]]
				c1 := guess_l[e[1]]
				for _, c2 := range colors {
					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[0]) * ctoi(c0)
					index += pow(6, e[1]) * ctoi(c1)
					index += pow(6, e[2]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[0]) * ctoi(c0)
					index += pow(6, e[2]) * ctoi(c1)
					index += pow(6, e[1]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[1]) * ctoi(c0)
					index += pow(6, e[0]) * ctoi(c1)
					index += pow(6, e[2]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[1]) * ctoi(c0)
					index += pow(6, e[2]) * ctoi(c1)
					index += pow(6, e[0]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[2]) * ctoi(c0)
					index += pow(6, e[0]) * ctoi(c1)
					index += pow(6, e[1]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[2]) * ctoi(c0)
					index += pow(6, e[1]) * ctoi(c1)
					index += pow(6, e[0]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}
				}
				c0 = guess_l[e[0]]
				c2 := guess_l[e[2]]
				for _, c1 := range colors {
					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[0]) * ctoi(c0)
					index += pow(6, e[1]) * ctoi(c1)
					index += pow(6, e[2]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[0]) * ctoi(c0)
					index += pow(6, e[2]) * ctoi(c1)
					index += pow(6, e[1]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[1]) * ctoi(c0)
					index += pow(6, e[0]) * ctoi(c1)
					index += pow(6, e[2]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[1]) * ctoi(c0)
					index += pow(6, e[2]) * ctoi(c1)
					index += pow(6, e[0]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[2]) * ctoi(c0)
					index += pow(6, e[0]) * ctoi(c1)
					index += pow(6, e[1]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[2]) * ctoi(c0)
					index += pow(6, e[1]) * ctoi(c1)
					index += pow(6, e[0]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}
				}
				c1 = guess_l[e[1]]
				c2 = guess_l[e[2]]
				for _, c0 := range colors {
					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[0]) * ctoi(c0)
					index += pow(6, e[1]) * ctoi(c1)
					index += pow(6, e[2]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[0]) * ctoi(c0)
					index += pow(6, e[2]) * ctoi(c1)
					index += pow(6, e[1]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[1]) * ctoi(c0)
					index += pow(6, e[0]) * ctoi(c1)
					index += pow(6, e[2]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[1]) * ctoi(c0)
					index += pow(6, e[2]) * ctoi(c1)
					index += pow(6, e[0]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[2]) * ctoi(c0)
					index += pow(6, e[0]) * ctoi(c1)
					index += pow(6, e[1]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}

					index = pow(6, ib) * ctoi(guess_l[ib])
					index += pow(6, e[2]) * ctoi(c0)
					index += pow(6, e[1]) * ctoi(c1)
					index += pow(6, e[0]) * ctoi(c2)
					if pool[index] != 0 {
						newpool[index] = 1
					}
				}

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
