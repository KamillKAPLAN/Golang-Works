package main

import "fmt"

// 001 Methods
type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " -ajan konuştu")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " -insan konuştu")
}

// 002 Interfaces
type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("Ben aradım...")
}

func main() {
	// 001 Methods
	sa1 := secretAgent{
		person: person{
			first: "Kamil",
			last:  "KAPLAN",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	// 002 Interfaces
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(p1)
	bar(sa1)
	bar(p1)
	p1.speak()
}
