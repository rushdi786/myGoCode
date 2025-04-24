package main

import "fmt"

func main() {
	var age int
	var name string
	var funFact string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("your name is", name)
	fmt.Scanln(&age)
	fmt.Println("your age is ", age)
	fmt.Print("Enter your fun fact: ")
	fmt.Scanln(&funFact)
	fmt.Println("your fun fact is ", funFact)
}
