package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[5])
	fmt.Println(x[1:3])

	fmt.Println("----------------")

	mySlice := make([]int, 0, 5)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fmt.Println("----------------")

	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}
