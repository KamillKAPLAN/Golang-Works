package main

import "fmt"

// 004 Example
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, " years old.")
}

func main() {
	// 001 Example
	x1 := foo()
	fmt.Println(x1)
	x2, x3 := bar()
	fmt.Println(x2, x3)

	// 002 Example
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := toplam(ii...)
	fmt.Println("Toplam: ", n)
	nn := toplam2(ii)
	fmt.Println("Toplam: ", nn)

	// 003 Example
	defer foo()
	bar()

	// 004 Example
	p1 := person{
		first: "Kamil",
		last:  "KAPLAN",
		age:   24,
	}
	p1.speak()
}

// 001 Example
func foo() int {
	fmt.Println("DEFER kullanımı")
	return 44
}

func bar() (int, string) {
	fmt.Println(1994, "Kamil KAPLAN doğdu.")
	return 1994, "Kamil KAPLAN doğdu."
}

// 002 Example
func toplam(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func toplam2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
