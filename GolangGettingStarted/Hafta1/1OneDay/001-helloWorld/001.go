package main

import "fmt"

func main() {
	fmt.Println("Merhaba ben Kamil KAPLAN")

	foo()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm foo")
}

func bar() {
	fmt.Println("I'm bar")
}

// control follow
//	- sequence
//	- loop; iterative
//	- conditional
