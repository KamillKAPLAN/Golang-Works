package main

import "fmt"

type person struct {
	name string
}

func main() {
	// 001 Example
	x := 54
	fmt.Println(x)
	fmt.Println(&x)

	// 002 Example
	p1 := person{
		name: "Kamil KAPLAN",
	}
	fmt.Println(p1)
	fmt.Println(&p1.name)
	changeMe(&p1)
	fmt.Println(p1)
}
func changeMe(p *person){
	p.name = "KAPLAN Kamil"
	//(*p).name = "asd"
}
