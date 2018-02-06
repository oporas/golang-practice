package main

import "fmt"

func main() {
	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Doe"
	student1[1] = "John"
	student1[2] = "100"
	student1[3] = "60"
	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Jooh"
	student2[1] = "Mage"
	student2[2] = "90"
	student2[3] = "40"
	records = append(records, student2)

	fmt.Println(records)
}
