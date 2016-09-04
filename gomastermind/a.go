package main

import (
	"fmt"
	//"sync"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd'}
	var n int = 4
	//var wg sync.WaitGroup
	c_chan := make(chan []byte)
	go Permute(c_chan, a, 0, n-1)
	Printer(c_chan)
}

func Permute(out chan<- []byte, a []byte, l, r int) {
	var i int
	if l == r {
		out <- a
		fmt.Println(string(a))

	} else {
		for i = l; i <= r; i++ {
			a[i], a[l] = a[l], a[i]
			Permute(out, a, l+1, r)
			a[i], a[l] = a[l], a[i]
		}
	}

}

func Printer(in <-chan []byte) {
	for i := range in {
		fmt.Println(string(i))
	}
}
