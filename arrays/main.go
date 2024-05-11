package main

import "fmt"

func main() {

	var fruits [4]string
	fruits[0] = "mango"
	fruits[1] = "apple"
	fruits[3] = "banana"
	fmt.Println("The fruits list is", fruits)
	fmt.Println("The fruits lenght is", len(fruits))

	var vegs = [3]string{"Patoto","Onions","Lady Finger"}

	fmt.Println("The veg list is",vegs)
}
