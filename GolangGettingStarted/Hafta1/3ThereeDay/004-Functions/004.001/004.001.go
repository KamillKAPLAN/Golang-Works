package main

import "fmt"

func main() {
	// 001 Function Syntax
	foo()
	bar("Kamil")
	s1 := woo("Para")
	fmt.Println(s1)
	x, y := mouse("Kamil", "Yasin")
	fmt.Println(x)
	fmt.Println(y)

	// 002 Variadic Parameter - Değişken Parametre
	//  -- s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//  -- fmt.Println(sum(s...))
	xa := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(xa)
	xi := sum2("Kamil")
	fmt.Println(xi)

	// 003
	defer dfr1()
	dfr2()

}

// func (r alıcı) tanımlayıcı(parameters) {
//		(return(s)) ...
// }
// fonksiyonumuzu parametrelerle tanımlarız - varsa
// fonksiyonumuzu çağırır ve argümanları iletiriz - herhangi birinde

// 001 Function Syntax
func foo() {
	fmt.Println("hello from foo")
}
func bar(s string) {
	fmt.Println("Hello,", s)
}
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` 'e "selam" söyledi.´`)
	b := false
	return a, b
}

// 002 Variadic Parameter - Değişken Parametre
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, val := range x {
		sum += val
		fmt.Println(val, " + ", sum)
	}
	return sum
}

func sum2(s string, x ...int) int {
	// tanımlanan fonksiyonda s değeri muhakkak olmalı fakat x olmasada olur
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, val := range x {
		sum += val
		fmt.Println(val, " + ", sum)
	}
	return sum
}

// 003 Defer - ertelee anahtar kelimesi
//Ertelenen çağrının argümanları derhal değerlendirilir, ancak işlev çağrısı çevreleyen işlev geri dönene kadar yürütülmez.

func dfr1() {
	fmt.Println("Defer çalıştı")
}
func dfr2() {
	fmt.Println("Başta çalıştı")
}
