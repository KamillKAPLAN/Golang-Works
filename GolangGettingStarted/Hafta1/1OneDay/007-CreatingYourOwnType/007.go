package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	// Kendi türünüzü oluşturma

	a = 42
	b = 43
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	// a = b
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// error: incompatible types in assignment
	// (cannot use type hotdog as type int)

	a = int(b)
	// tip dönüşümü yaparsak atama işlemini yaparız.
}
