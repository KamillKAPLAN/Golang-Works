package main

import "fmt"

var y string
var z int

func main() {
	// belirli bir türden bir değişken ilan etmek
	// ve sonra değişkene bu tür bir değer atayın

	fmt.Println("Y'nin değeri: ", y, "son")
	fmt.Printf("%T\n", y)

	y = "Kamil KAPLAN"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
