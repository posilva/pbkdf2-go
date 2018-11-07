package main

import (
	"fmt"
	"log"
)

func main() {
	code, err := generateAccessCode()
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Code: %v \n", code)
	fmt.Printf("Code Hashed: %v", HashPassword(code))

}
