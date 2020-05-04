package main

import "fmt"

var xx int8 = -128

func main() {
	// 001. Bool type
	a := 7
	b := 42
	fmt.Println(a == b)

	// 002. Numeric type
	x := 42
	y := 42.34534
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(xx)
	fmt.Printf("%T\n", xx)

	// 003. String type
	ss := "Merhaba ben Kamil KAPLAN"
	fmt.Println(ss)
	fmt.Printf("%T\n", ss)

	bbs := []byte(ss)
	fmt.Println(bbs)
	fmt.Printf("%T\n", bbs)

	// 004.Numeral(Sayısal) system
	s := "H"
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println("Byte karşılığı: ", bs)
	n := bs[0]
	fmt.Println("b[0] daki değer: ", n)

	// 005.Constants ( Sabitler)
	const a1 int = 42
	const b1 = 42.78
	const c1 = "Kamil KAPLAN"
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", c1)

	const (
		ortak1 int    = 44
		ortak2        = 44.44
		ortak3 string = "Malatya"
		ortak4 bool   = true
	)
	fmt.Println(ortak1)
	fmt.Println(ortak2)
	fmt.Println(ortak3)
	fmt.Println(ortak4)
}
