package main

import "fmt"

func main() {
	n := average(42, 56, 87, 21, 45, 48)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	// total := 0.0
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
