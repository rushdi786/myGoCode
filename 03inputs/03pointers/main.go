package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Value of pointer: ", ptr)

	myNum := 23

	var myPointer = &myNum //ref the location of the value of mynum
	fmt.Println("Value of pointer: ", myPointer)
	fmt.Println("Value of pointer: ", *myPointer)

	myPointer = *myPointer * 2
	fmt.Println("Value of pointer: ", myPointer)
	fmt.Println("Value of pointer: ", *myPointer)

}
