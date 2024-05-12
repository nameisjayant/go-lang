package main

import "fmt"

func main()  {
	
	data := User{"Jayant","25",true}
	fmt.Println(data)
	fmt.Printf("The details are %+v\n",data)
	fmt.Println("The name is : ",data.Name)
}

type User struct {
	Name string
	Age string
	IsVerfied  bool
}