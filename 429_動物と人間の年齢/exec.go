package main

import (
	"fmt"
)

func main() {
	fmt.Println(qualifiedForInsurance(dogToHuman, 7))
	fmt.Println(qualifiedForInsurance(dogToHuman, 8))
	fmt.Println(qualifiedForInsurance(catToHuman, 11))
	fmt.Println(qualifiedForInsurance(catToHuman, 12))
}

type AgeConverter func(int) int

func qualifiedForInsurance(callable AgeConverter, age int) bool {
	if callable(age) <= 60 {
		return true
	} else {
		return false
	}
}

func dogToHuman(dogAge int) int {
	return 20 + (dogAge-2)*7
}

func catToHuman(catAge int) int {
	return 24 + (catAge-2)*4
}
