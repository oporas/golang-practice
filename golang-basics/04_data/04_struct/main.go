package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First        string
	LicenseToKil bool
}

func (p Person) fullName() string { // (p Person) = receiver, so function/method is added to Person struct
	return p.First + " " + p.Last
}

func main() {
	jooh := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:        "Mah name is bond",
		LicenseToKil: true,
	}

	fmt.Println(jooh)
	fmt.Println(jooh.Age)          // DoubleZero gets the .Age from Person
	fmt.Println(jooh.First)        // .First overriden by DoubleZero
	fmt.Println(jooh.Person.First) // .First still accessible from Person

	p1 := Person{"James", "Bond", 30}
	p2 := Person{"Miss", "Moneypenny", 20}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
