package main

import "fmt"

func main() {
	data := []float64{42, 56, 87, 21, 45, 48}
	n := averageRange(data...)
	fmt.Println(n)
	n = averageFloat64(data)
	fmt.Println(n)
}

func averageRange(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	// total := 0.0
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func averageFloat64(sf []float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	// total := 0.0
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
