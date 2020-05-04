package main

import (
	"fmt"
)

func main() {
	var answer1, answer2, answer3 string
	fmt.Print("Name: ")
	_, err2 := fmt.Scan(&answer1)
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("Fav Food: ")
	_, err2 = fmt.Scan(&answer2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("Fav Sport: ")
	_, err2 = fmt.Scan(&answer3)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(answer1, answer2, answer3)
}
