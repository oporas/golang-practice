package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	var y [58]int
	fmt.Println(y)
	fmt.Println(y[12])
	y[12] = 666
	fmt.Println(y[12])

	for i, v := range y {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 20 {
			break
		}
	}
}
