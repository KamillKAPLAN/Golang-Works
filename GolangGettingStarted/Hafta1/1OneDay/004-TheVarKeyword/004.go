package main

import "fmt"

// değişkenini "y" olarak bildir
// 43 değerini atayın
// bildir & ata = başlatma
var y = 43

// tanımlayıcı ile bir değişken olduğunu bildirir
// ve "z" tanımlayıcılı değişken, int türündedir
// int türünün sıfır değerini "z" ye atar
// booleans için false, interger'lar için 0, float'lar için 0.0, string'ler için ""
// pointers, functions, slices, chanels ve maps için nil
var z int

func main() {
	// kısa bildirim operatörü
	// Bir değişkeni tanımlayın ve bir değer atayın (belirli bir türde)
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again: ", y)
}
