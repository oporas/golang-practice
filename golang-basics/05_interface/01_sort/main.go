package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	//Same with strings
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	fmt.Println(s)

	//Short version
	s2 := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s2)
	sort.Strings(s2)
	fmt.Println(s2)

	//Reverse
	s3 := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s3)
	sort.Sort(sort.Reverse(sort.StringSlice(s3)))
	fmt.Println(s3)

	//int
	n := []int{3, 1, 4, 8, 5, 2, 3, 8, 25, 12, 65}
	fmt.Println(n)
	// sort.Sort(sort.IntSlice(n))
	sort.Ints(n)
	fmt.Println(n)

}
