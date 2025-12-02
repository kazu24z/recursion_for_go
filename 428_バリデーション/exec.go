package main

import (
	"fmt"
)

func main() {
	fmt.Println(emailValidation(doesNotStartWithAt, "@gmail.com"))
	fmt.Println(emailValidation(doesNotStartWithAt, "kkk@gmail.com"))
	fmt.Println(emailValidation(doesNotHaveSpace, "Hello world"))
	fmt.Println(emailValidation(doesNotHaveSpace, "Helloworld"))
	fmt.Println(emailValidation(hasUppercaseAndLowercase, "hello world"))
	fmt.Println(emailValidation(hasUppercaseAndLowercase, "HELLO WORLD"))
	fmt.Println(emailValidation(hasUppercaseAndLowercase, "Hello world"))
}

type Validator func(string) bool

func emailValidation(v Validator, s string) string {
	if v(s) {
		return "Email is correct."
	} else {
		return "Email is not correct."
	}
}

func doesNotStartWithAt(email string) bool {
	if len(email) > 0 && email[0] != '@' {
		return true
	} else {
		return false
	}
}

func doesNotHaveSpace(email string) bool {
	for i := 0; i < len(email); i++ {
		if email[i] == ' ' {
			return false
		}
	}
	return true
}

func hasUppercaseAndLowercase(email string) bool {
	hasUpper, hasLower := false, false

	for i := 0; i < len(email); i++ {
		if 'A' <= email[i] && email[i] <= 'Z' {
			hasUpper = true
		}

		if 'a' <= email[i] && email[i] <= 'z' {
			hasLower = true
		}

		if hasUpper && hasLower {
			return true
		}
	}
	return false
}
