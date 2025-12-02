package main

import (
	"fmt"
)

func main() {
	fmt.Println(summation(isOdd, 3))
	fmt.Println(summation(isOdd, 10))
	fmt.Println(summation(isOdd, 25))
	fmt.Println(summation(isMultipleOf3Or5, 3))
	fmt.Println(summation(isMultipleOf3Or5, 10))
	fmt.Println(summation(isMultipleOf3Or5, 100))
	fmt.Println(summation(isPrime, 2))
	fmt.Println(summation(isPrime, 10))
	fmt.Println(summation(isPrime, 100))

}

type Condition func(int) bool

func summation(f Condition, n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if f(i) {
			sum = sum + i
		}
	}
	return sum
}

func isOdd(x int) bool {
	if x%2 != 0 {
		return true
	}
	return false
}

func isMultipleOf3Or5(x int) bool {
	if x%3 == 0 || x%5 == 0 {
		return true
	}
	return false
}

// これ、xの平方根までのチェックいい
func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
