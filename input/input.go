package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name?")

	// common ok | common err syntax
	input,_ := reader.ReadString('\n');
	print("Your name is ",input)
}