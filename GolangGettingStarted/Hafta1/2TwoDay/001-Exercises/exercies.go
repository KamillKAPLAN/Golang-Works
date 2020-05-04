package main

import "fmt"

var a int
var b string
var c bool

var aa int = 12
var bb string = "Kamil KAPLAN"
var cc bool = false

type deneme int
var ax deneme

type hotdog int 
var hotX hotdog
var hotY int

func main() {
	// 001 Example
	x := 42
	y := "Kamil KAPLAN"
	z := true
	fmt.Println(x, " ", y, " ", z)

	// 002 Example
	fmt.Println(a, " ", b, " ", c)

	// 003 Example
	fmt.Println(aa, " ", bb, " ", cc)

	// 004 Example
	fmt.Println(ax)
	fmt.Printf("%T\n", ax)
	ax = 35
	fmt.Println(ax)

	// 005 Example
	fmt.Println(hotX)
	fmt.Printf("%T\n", hotX)
	hotX = 44
	fmt.Println(hotX)
	hotY = int(hotX)
	fmt.Println(hotY)
	fmt.Printf("%T\n", hotY)
}
