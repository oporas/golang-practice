package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
	Score int `json:"wisdom score"` //TAG
}

func main() {
	//Unmarshal
	var p1 person
	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20, "wisdom score":40}`)
	json.Unmarshal(bs, &p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Println(p1.Score)

	//Encode
	p2 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p2) // Writed to Stdout/terminal

	// Decode
	var p3 person
	reader := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20, "wisdom score":40}`)
	json.NewDecoder(reader).Decode(&p3)

	fmt.Println(p3.First)
	fmt.Println(p3.Last)
	fmt.Println(p3.Age)
	fmt.Println(p3.Score)

}
