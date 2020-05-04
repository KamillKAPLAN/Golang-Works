package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type gomuluStruct struct {
	person
	onay bool
}

func main() {
	// 001 Struct
	p1 := person{
		first: "Kamil",
		last:  "KAPLAN",
		age:   25,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)

	// 002 Gömülü Struct
	gml := gomuluStruct{
		person: person{
			first: "Gomulu Ad",
			last:  "Gomulu Soyad",
			age:   24,
		},
		onay: true,
	}
	fmt.Println(gml.first)
}
