package main

import "fmt"

// const variable always start with capital letter...
const LoginToken  string  = "gdhjdkf"

func main() {
	var name string = "Jayant"
	fmt.Println(name)
	fmt.Printf("The type of %s is %T \n", name, name)

	var number int = 13
	println(number)
	var isLogged bool = true
	println(isLogged)

	var decimal float64 = 2.34
	println(decimal)

	// implicit type

	var website = "nameisjayant.com"
	fmt.Println(website)

	// no var style

	number_1 := 4000
	fmt.Println(number_1)

}
