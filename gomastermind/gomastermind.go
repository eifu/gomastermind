package gomastermind

import (
	"fmt"
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

func split(s string) []byte {
	a := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = s[i]
	}
	return a
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

func Finder(guess_l []byte, score_l, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors []byte = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
	var c0, c1, c2, c3 byte
	var b1 [4][]int = [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	//	var b2 [6][]int = [6][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}}
	//var b3 [4][]int = [4][]int{[]int{1}, []int{2}, []int{3}, []int{4}}
	var index int
	if score_l[0] == 0 {
		// b=0

		if score_l[1] == 0 {
			// w=0
			for i := 0; i < 4; i++ {
				colors = elimColor(guess_l[i], colors)
			}
			for _, c0 = range colors {
				for _, c1 = range colors {
					for _, c2 = range colors {
						for _, c3 = range colors {
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
			return newpool
		} else if score_l[1] == 1 {
			// w=1
			// pick a color from guess_l, then remove all colors in guess_l from colors
			for c0pos := 0; c0pos < 4; c0pos++ {
				c0 = guess_l[c0pos] // pick a color from guess_l

				combos := [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}

				// remove the rest of three colors from the potential colors
				colors = elimColor(guess_l[combos[c0pos][0]], colors)
				colors = elimColor(guess_l[combos[c0pos][1]], colors)
				colors = elimColor(guess_l[combos[c0pos][2]], colors)

				for _, c1 = range colors { // c1, c2 and c3 are topological
					for _, c2 = range colors {
						for _, c3 = range colors {
							for _, dst := range combos[c0pos] {
								if c0 != guess_l[dst] &&
									c1 != guess_l[combos[dst][0]] &&
									c2 != guess_l[combos[dst][1]] &&
									c3 != guess_l[combos[dst][2]] {

									index = pow(6, dst) * ctoi(c0)
									index += pow(6, combos[dst][0]) * ctoi(c1)
									index += pow(6, combos[dst][1]) * ctoi(c2)
									index += pow(6, combos[dst][2]) * ctoi(c3)
									if pool[index] != 0 {
										newpool[index] = 1
									}
								}
							}
						}
					}
				}
			}
			return newpool

		} else if score_l[1] == 2 {
			// w=2
			combos := [...][]int{[]int{0, 1}, []int{1, 0},
				[]int{0, 2}, []int{2, 0},
				[]int{0, 3}, []int{3, 0},
				[]int{1, 2}, []int{2, 1},
				[]int{1, 3}, []int{3, 1},
				[]int{2, 3}, []int{3, 2}}

			for isrc, src := range combos {

				colors = elimColor(guess_l[combos[11-isrc][0]], colors)
				colors = elimColor(guess_l[combos[11-isrc][1]], colors)

				c0 = guess_l[src[0]]
				c1 = guess_l[src[1]]

				for _, c2 = range colors {
					for _, c3 = range colors {
						// c2 and c3 are topological

						for idst, dst := range combos {

							if c0 != guess_l[dst[0]] &&
								c1 != guess_l[dst[1]] &&
								c2 != guess_l[combos[11-idst][0]] &&
								c3 != guess_l[combos[11-idst][1]] {
								index = pow(6, dst[0]) * ctoi(c0)
								index += pow(6, dst[1]) * ctoi(c1)
								index += pow(6, combos[11-idst][0]) * ctoi(c2)
								index += pow(6, combos[11-idst][1]) * ctoi(c3)
								if pool[index] != 0 {
									newpool[index] = 1
								}
							}
						}
					}
				}
			}
			return newpool
		} else if score_l[1] == 3 {
			// w=3

			combos := [...][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
			// possible combination to take 3 numbers from 0, 1, 2, 3
			permus := [...][]int{[]int{0, 1, 2}, []int{0, 2, 1}, []int{1, 0, 2}, []int{1, 2, 0}, []int{2, 0, 1}, []int{2, 1, 0}}
			// possible permutation to choose 3 numbers

			for isrc, src := range combos {
				colors = elimColor(guess_l[isrc], colors)
				for _, permu := range permus {
					c0 = guess_l[src[permu[0]]]
					c1 = guess_l[src[permu[1]]]
					c2 = guess_l[src[permu[2]]]

					for _, c3 = range colors {
						for idst, dst := range combos {
							for _, permu := range permus {
								if c0 != guess_l[dst[permu[0]]] &&
									c1 != guess_l[dst[permu[1]]] &&
									c2 != guess_l[dst[permu[2]]] &&
									c3 != guess_l[idst] {

									index = pow(6, dst[permu[0]]) * ctoi(c0)
									index += pow(6, dst[permu[1]]) * ctoi(c1)
									index += pow(6, dst[permu[2]]) * ctoi(c2)
									index += pow(6, idst) * ctoi(c3)

									if pool[index] != 0 {
										newpool[index] = 1
									}
								}
							}
						}
					}
				}
			}
			return newpool
		} else if score_l[1] == 4 {
			// w=4
			combos := [...][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
			// possible combination to take 3 numbers from 0, 1, 2, 3
			permus := [...][]int{[]int{0, 1, 2}, []int{0, 2, 1}, []int{1, 0, 2}, []int{1, 2, 0}, []int{2, 0, 1}, []int{2, 1, 0}}
			// possible permutation to choose 3 numbers
			for idst, dst := range combos {
				for _, permu := range permus {
					if guess_l[0] != guess_l[dst[permu[0]]] &&
						guess_l[1] != guess_l[dst[permu[1]]] &&
						guess_l[2] != guess_l[dst[permu[2]]] &&
						guess_l[3] != guess_l[idst] {
						index = pow(6, dst[permu[0]]) * ctoi(guess_l[0])
						index += pow(6, dst[permu[1]]) * ctoi(guess_l[1])
						index += pow(6, dst[permu[2]]) * ctoi(guess_l[2])
						index += pow(6, idst) * ctoi(guess_l[3])
						if pool[index] != 0 {
							newpool[index] = 1
						}
					}
				}
			}
			return newpool
		}

	} else if score_l[0] == 1 {
		// b=1
		if score_l[1] == 0 {
			// w=0
			for ib, belects := range b1 {
				c0 = guess_l[ib]

				for _, belect := range belects {
					colors = elimColor(guess_l[belect], colors)
				}
				for _, c1 := range colors { // 1st
					for _, c2 := range colors { // 2nd
						for _, c3 := range colors { // 3rd
							index = pow(6, ib) * ctoi(c0)
							index += pow(6, belects[0]) * ctoi(c1)
							index += pow(6, belects[1]) * ctoi(c2)
							index += pow(6, belects[2]) * ctoi(c3)
							if pool[index] != 0 {
								newpool[index] = 1
							}
						}
					}
				}

			}
		} else if score_l[1] == 1 {
			// w=1
			combos := [...][]int{[]int{1, 2}, []int{0, 2}, []int{0, 1}}
			for ib, belects := range b1 {
				colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
				for isrc, src := range combos {
					c0 = guess_l[ib]            // c0 is the right color in the right position
					c1 = guess_l[belects[isrc]] // c1 is the fixed color, not in the position
					colors = elimColor(guess_l[belects[src[0]]], colors)
					colors = elimColor(guess_l[belects[src[1]]], colors)
					for _, c2 := range colors { // c2 and c3 are topologic
						for _, c3 := range colors {
							for idst, dst := range combos {
								if c1 != guess_l[belects[idst]] &&
									c2 != guess_l[belects[dst[0]]] &&
									c3 != guess_l[belects[dst[1]]] {
									index = pow(6, ib) * ctoi(c0)
									index += pow(6, belects[idst]) * ctoi(c1)
									index += pow(6, belects[dst[0]]) * ctoi(c2)
									index += pow(6, belects[dst[1]]) * ctoi(c3)
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
			// w=2 01, 02, 12

			combos := [...][]int{[]int{1, 2}, []int{0, 2}, []int{0, 1}}
			for ib, belects := range b1 {
				for isrc, src := range combos {

					c0 := guess_l[belects[src[0]]]
					c1 := guess_l[belects[src[1]]]
					colors = elimColor(guess_l[belects[isrc]], colors)
					for _, c2 := range colors {
						for idst, dst := range combos {
							if c0 != guess_l[belects[dst[0]]] &&
								c1 != guess_l[belects[dst[1]]] &&
								c2 != guess_l[belects[idst]] {
								index = pow(6, ib) * ctoi(guess_l[ib])
								index += pow(6, belects[dst[0]]) * ctoi(c0)
								index += pow(6, belects[dst[1]]) * ctoi(c1)
								index += pow(6, belects[idst]) * ctoi(c2)
								if pool[index] != 0 {
									newpool[index] = 1
								}
							}

							if c0 != guess_l[belects[dst[1]]] &&
								c1 != guess_l[belects[dst[0]]] &&
								c2 != guess_l[belects[idst]] {
								index = pow(6, ib) * ctoi(guess_l[ib])
								index += pow(6, belects[dst[1]]) * ctoi(c0)
								index += pow(6, belects[dst[0]]) * ctoi(c1)
								index += pow(6, belects[idst]) * ctoi(c2)
								if pool[index] != 0 {
									newpool[index] = 1
								}
							}
						}
					}
				}
			}

		}

		return newpool
	} else if score_l[0] == 2 {

	} else if score_l[0] == 3 {

	} else {

	}

	return newpool
}
