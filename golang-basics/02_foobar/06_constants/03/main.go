package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

const (
	_ = iota      // 0
	g = iota * 10 // 10
	h = iota * 10 // 20
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
