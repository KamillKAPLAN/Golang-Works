package main

import "fmt"

func main() {
	// 001 Lessons
	a := 44
	fmt.Println(a)
	// a değerinin bellekteki adresini verir bize
	fmt.Println(&a)

	// a değerinin veri tipini verir
	fmt.Printf("%T\n", a)
	// a değerinin bellek tipini verir
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	// a değerinin belleğe kaydedilmiş adresteki değerini verir
	fmt.Println(*b)

	// 002 Lessons
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", x)
	fmt.Println("x after", x)
}

// 002 Lessons
func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
