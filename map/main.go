package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's learn about maps in go")
	languages := make(map[string]string)
	languages["kt"] = "Kotlin"
	languages["py"] = "python"
	languages["rb"] = "Ruby"
	languages["js"] = "Javascript"

	fmt.Println("The languages are:", languages)

	fmt.Println("Kotlin value : ", languages["kt"])

	delete(languages, "rb")
	fmt.Println("The languages are:", languages)

	//loops

	for key, value := range languages {
		fmt.Printf("The key %v , value is %v \n",key,value)
	}
}
