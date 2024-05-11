package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println("The value of pointer is",ptr)

	var number = 23
	var numPtr = &number

	fmt.Println("The address is ",numPtr)
	fmt.Println("The value is ",*numPtr)

	*numPtr = *numPtr + 2
	fmt.Println("The new value is",number)
}
