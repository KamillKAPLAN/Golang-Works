package main

import "fmt"

var y = 42

// değişken "z" tanımlayıcısı ile bildirildi
// string türünde
// ve "Kamil KAPLAN" değerini atadım
var z string = "Kamil KAPLAN"
var a = `Kamil söyledi, "Hava nasıl" `

// bu statik bir programlama dilidir
// belirli bir türün değerini tutmak için bir değişken bildirilir
// dinamik bir programlama dili değil

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// z = 43
	// fmt.Println(z)
}
