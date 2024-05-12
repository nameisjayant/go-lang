package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Learn about Slices in Go lang...")
	var fruits = []string{"Apple", "Banana", "Orange"}
	fmt.Println("The fruits are", fruits)
	fruits = append(fruits, "Mango", "Tomato")
	fmt.Println("Update fruits are : ", fruits)

	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	// make keyword to allocate memory

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 354
	highscores[2] = 897
	highscores[3] = 455

	fmt.Println(highscores)
	highscores = append(highscores, 222, 777)
	fmt.Println(highscores)
	// sort the slices
	sort.Ints(highscores)
	fmt.Println(highscores)

	// remove  slices  based on index

	var courses = []string{"React js", "Kotlin", "Swift", "Go", "Java"}
	indexToRemove := 2
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println(courses)
}
