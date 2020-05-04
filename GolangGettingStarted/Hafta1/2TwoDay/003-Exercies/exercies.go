package main

import "fmt"

func main() {
	// 001 Binary convert
	a := 5
	fmt.Printf("%d\t%b\n", a, a)

	// 002 example
	b := (42 == 42)
	c := (42 <= 43)
	d := (42 >= 43)
	e := (42 != 43)
	f := (42 < 43)
	g := (42 > 43)
	fmt.Println(b, c, d, e, f, g)

	// 003 const example
	const (
		cnst1     = 42
		cnst2 int = 43
	)
	fmt.Println(cnst1, cnst2)

}
