package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(greet(1, "John", loud))
	fmt.Println(greet(2, "John", quiet))
	fmt.Println(greet(13, "John", reverse))
	fmt.Println(greet(19, "John", repeat))
	fmt.Println(greet(13, "Leslie Emmanuel Beadon", loud))

	fmt.Println(greet(19, "Leslie Emmanuel Beadon", quiet))

	fmt.Println(greet(5, "Leslie Emmanuel Beadon", reverse))

	fmt.Println(greet(1, "Leslie Emmanuel Beadon", repeat))
}

type Operation func(string) string

func greet(time int, name string, f Operation) string {
	greetStr := ""
	if 0 <= time && time < 12 {
		greetStr = "Good Morning"
	} else if 12 <= time && time < 18 {
		greetStr = "Good Afternoon"
	} else {
		greetStr = "Good Evening"
	}

	return greetStr + " " + f(name)
}

func loud(text string) string {
	return strings.ToUpper(text)
}

func quiet(text string) string {
	return strings.ToLower(text)
}

func reverse(text string) string {
	reversed := []byte("")
	for i := 0; i < len(text); i++ {
		reversed = append(reversed, text[len(text)-1-i])
	}
	return string(reversed)
}

func repeat(text string) string {
	return text + " " + text
}
