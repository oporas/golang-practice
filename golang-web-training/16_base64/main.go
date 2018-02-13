package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Testing testing jooh testing"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("It's me mario", err)
	}
	fmt.Println(string(bs))
}
