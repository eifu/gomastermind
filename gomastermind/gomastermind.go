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

func permute(a []byte, l, r int) {
	var i int
	if l == r {
		fmt.Printf("%v\n", string(a))
	} else {
		for i = l; i <= r; i++ {
			a[i], a[l] = a[l], a[i]
			permute(a, l+1, r)
			a[i], a[l] = a[l], a[i]
		}
	}
}

func b0w0(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors []byte = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
	var index int
	for i := 0; i < 4; i++ {
		// remove all colors in guess
		colors = elimColor(guess_l[i], colors)
	}
	var c0, c1, c2, c3 byte
	for _, c0 = range colors {
		for _, c1 = range colors {
			for _, c2 = range colors {
				for _, c3 = range colors {
					// c0, c1, c2, c3 are topological
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
}

func b0w1(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors []byte = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
	var c0, c1, c2, c3 byte
	var index int
	combos := [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	// pick a color from guess_l, then remove the rest of three colors in guess_l from colors
	for c0pos, elective := range combos {
		c0 = guess_l[c0pos] // pick a color from guess_l
		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		// remove the rest of three colors from the potential colors
		colors = elimColor(guess_l[elective[0]], colors)
		colors = elimColor(guess_l[elective[1]], colors)
		colors = elimColor(guess_l[elective[2]], colors)

		for _, c1 = range colors { // c1, c2 and c3 are topological
			for _, c2 = range colors {
				for _, c3 = range colors {

					// TODO: need to fix this. there is something unclear about this loop

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
}

func b0w2(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors []byte = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
	var c0, c1, c2, c3 byte
	var index int
	combos := [...][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 2}, []int{1, 3}, []int{2, 3}}
	permus := [...][]int{[]int{0, 1}, []int{1, 0}}

	for isrc, src := range combos {

		// c0 and c1 are topological
		c0 = guess_l[src[0]]
		c1 = guess_l[src[1]]

		// remove the rest of two colors
		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		colors = elimColor(guess_l[combos[5-isrc][0]], colors)
		colors = elimColor(guess_l[combos[5-isrc][1]], colors)

		// c2 and c3 are topological
		for _, c2 = range colors {
			for _, c3 = range colors {
				// TODO: make function that gives permutation of 4 numbers

				for _, permu := range permus {
					if c0 != guess_l[combos[5-isrc][permu[0]]] &&
						c1 != guess_l[combos[5-isrc][permu[1]]] &&
						c2 != guess_l[src[0]] &&
						c3 != guess_l[src[1]] {
						index = pow(6, combos[5-isrc][permu[0]]) * ctoi(c0)
						index += pow(6, combos[5-isrc][permu[1]]) * ctoi(c1)
						index += pow(6, src[0]) * ctoi(c2)
						index += pow(6, src[1]) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}
					}
				}
			}
		}
	}
	return newpool
}

func b0w3(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var colors []byte
	var c0, c1, c2, c3 byte
	var index int
	combos := [...][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	// possible combination to take 3 numbers from 0, 1, 2, 3
	permus := [...][]int{[]int{0, 1, 2}, []int{0, 2, 1}, []int{1, 0, 2}, []int{1, 2, 0}, []int{2, 0, 1}, []int{2, 1, 0}}
	// possible permutation to choose 3 numbers

	for isrc, src := range combos {
		// src is three indices that are right colors but not in the right position
		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		// remove 1 color from the colors
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
}

func b0w4(guess_l []byte, pool []int) []int {
	var index int
	newpool := make([]int, 6*6*6*6)
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

func b1w0(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var b1 [4][]int = [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	var colors []byte
	var c0, c1, c2, c3 byte
	var index int
	for ib, belects := range b1 {
		c0 = guess_l[ib]
		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		for _, belect := range belects {
			colors = elimColor(guess_l[belect], colors)
		}
		for _, c1 = range colors { // 1st
			for _, c2 = range colors { // 2nd
				for _, c3 = range colors { // 3rd
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
	return newpool
}

func b1w1(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var b1 [4][]int = [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	var index int
	var c0, c1, c2, c3 byte
	var colors []byte
	combos := [...][]int{[]int{1, 2}, []int{0, 2}, []int{0, 1}}
	for ib, belects := range b1 {

		for isrc, src := range combos {
			c0 = guess_l[ib]            // c0 is the right color in the right position
			c1 = guess_l[belects[isrc]] // c1 is the fixed color, not in the position
			colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
			colors = elimColor(guess_l[belects[src[0]]], colors)
			colors = elimColor(guess_l[belects[src[1]]], colors)
			for _, c2 = range colors { // c2 and c3 are topologic
				for _, c3 = range colors {
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
	return newpool
}

func b1w2(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var b1 [4][]int = [4][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	var c0, c1, c2, c3 byte
	var colors []byte
	var index int
	combos := [...][]int{[]int{1, 2}, []int{0, 2}, []int{0, 1}}
	for ib, belects := range b1 {
		// c0 is right color in right position
		c0 = guess_l[ib]

		for isrc, src := range combos {
			// c1, c2 are right color, but un wrong position
			c1 = guess_l[belects[src[0]]]
			c2 = guess_l[belects[src[1]]]

			// remove 1 color out of rest of 3 colors.
			colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
			colors = elimColor(guess_l[belects[isrc]], colors)

			for _, c3 = range colors {
				for idst, dst := range combos {
					if c1 != guess_l[belects[dst[0]]] &&
						c2 != guess_l[belects[dst[1]]] &&
						c3 != guess_l[belects[idst]] {
						index = pow(6, ib) * ctoi(c0)
						index += pow(6, belects[dst[0]]) * ctoi(c1)
						index += pow(6, belects[dst[1]]) * ctoi(c2)
						index += pow(6, belects[idst]) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}
					}

					if c1 != guess_l[belects[dst[1]]] &&
						c2 != guess_l[belects[dst[0]]] &&
						c3 != guess_l[belects[idst]] {
						index = pow(6, ib) * ctoi(c0)
						index += pow(6, belects[dst[1]]) * ctoi(c1)
						index += pow(6, belects[dst[0]]) * ctoi(c2)
						index += pow(6, belects[idst]) * ctoi(c3)
						if pool[index] != 0 {
							newpool[index] = 1
						}
					}
				}
			}
		}
	}

	return newpool

}

func b2w0(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var c0, c1, c2, c3 byte
	var colors []byte
	var index int
	combos := [...][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 2}, []int{1, 3}, []int{2, 3}}
	permus := [...][]int{[]int{0, 1}, []int{1, 0}}

	for isrc, src := range combos {

		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		colors = elimColor(guess_l[combos[5-isrc][0]], colors)
		colors = elimColor(guess_l[combos[5-isrc][1]], colors)

		for _, permu := range permus {

			c0 = guess_l[src[permu[0]]]
			c1 = guess_l[src[permu[1]]]

			for _, c2 = range colors {
				for _, c3 = range colors {

					if c2 != guess_l[combos[5-isrc][0]] &&
						c3 != guess_l[combos[5-isrc][1]] {

						index = pow(6, src[permu[0]]) * ctoi(c0)
						index += pow(6, src[permu[1]]) * ctoi(c1)
						index += pow(6, combos[5-isrc][0]) * ctoi(c2)
						index += pow(6, combos[5-isrc][1]) * ctoi(c3)

						if pool[index] != 0 {
							newpool[index] = 1
						}
					}

				}
			}
		}
	}
	return newpool
}

func b2w1(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	combos := [...][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 2}, []int{1, 3}, []int{2, 3}}
	var c0, c1, c2, c3 byte
	var colors []byte
	var index int
	for isrc, src := range combos {

		c0 = guess_l[src[0]]
		c1 = guess_l[src[1]]
		c2 = guess_l[combos[5-isrc][0]]
		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		colors = elimColor(guess_l[combos[5-isrc][1]], colors)

		for _, c3 = range colors {
			if c2 != guess_l[combos[5-isrc][1]] &&
				c3 != guess_l[combos[5-isrc][0]] {

				index = pow(6, src[0]) * ctoi(c0)
				index += pow(6, src[1]) * ctoi(c1)
				index += pow(6, combos[5-isrc][1]) * ctoi(c2)
				index += pow(6, combos[5-isrc][0]) * ctoi(c3)

				if pool[index] != 0 {
					newpool[index] = 1
				}
			}
		}

		c2 = guess_l[combos[5-isrc][1]]
		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		colors = elimColor(guess_l[combos[5-isrc][0]], colors)

		for _, c3 = range colors {
			if c2 != guess_l[combos[5-isrc][0]] &&
				c3 != guess_l[combos[5-isrc][1]] {

				index = pow(6, src[0]) * ctoi(c0)
				index += pow(6, src[1]) * ctoi(c1)
				index += pow(6, combos[5-isrc][0]) * ctoi(c2)
				index += pow(6, combos[5-isrc][1]) * ctoi(c3)

				if pool[index] != 0 {
					newpool[index] = 1
				}
			}

		}
	}
	return newpool
}

func b3w0(guess_l []byte, pool []int) []int {
	newpool := make([]int, 6*6*6*6)
	var c0, c1, c2, c3 byte
	var colors []byte
	var index int
	combos := [...][]int{[]int{1, 2, 3}, []int{0, 2, 3}, []int{0, 1, 3}, []int{0, 1, 2}}
	for isrc, src := range combos {
		c0 = guess_l[src[0]]
		c1 = guess_l[src[1]]
		c2 = guess_l[src[2]]

		colors = []byte{'R', 'W', 'Y', 'G', 'U', 'K'}
		colors = elimColor(guess_l[isrc], colors)
		for _, c3 = range colors {
			index = pow(6, src[0]) * ctoi(c0)
			index += pow(6, src[1]) * ctoi(c1)
			index += pow(6, src[2]) * ctoi(c2)
			index += pow(6, isrc) * ctoi(c3)
			if pool[index] != 0 {
				newpool[index] = 1
			}
		}

	}
	return newpool
}

func Finder(guess_l []byte, score_l, pool []int) []int {

	if score_l[0] == 0 && score_l[1] == 0 {
		// black 0, white 0
		return b0w0(guess_l, pool)
	} else if score_l[0] == 0 && score_l[1] == 1 {
		// black 0, white 1
		return b0w1(guess_l, pool)
	} else if score_l[0] == 0 && score_l[1] == 2 {
		// black 0, white 2
		return b0w2(guess_l, pool)
	} else if score_l[0] == 0 && score_l[1] == 3 {
		// black 0, white 3
		return b0w3(guess_l, pool)
	} else if score_l[0] == 0 && score_l[1] == 4 {
		// black 0, white 4
		return b0w4(guess_l, pool)
	} else if score_l[0] == 1 && score_l[1] == 0 {
		// black 1, white 0
		return b1w0(guess_l, pool)
	} else if score_l[0] == 1 && score_l[1] == 1 {
		// black 1, white 1
		return b1w1(guess_l, pool)
	} else if score_l[0] == 1 && score_l[1] == 2 {
		// black 1, white 2
		return b1w2(guess_l, pool)
	} else if score_l[0] == 2 && score_l[1] == 0 {
		// black 2, white 0
		return b2w0(guess_l, pool)
	} else if score_l[0] == 2 && score_l[1] == 1 {
		return b2w1(guess_l, pool)
	} else if score_l[0] == 3 {
		// b=3
		return b3w0(guess_l, pool)
	}
	return pool
}
