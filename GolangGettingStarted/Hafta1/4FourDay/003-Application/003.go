package main

import (
	"fmt"
	"encoding/json"
	"sort"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	// 002 JSON marshal
	p1 := person{
		First: "Kamil",
		Last:  "KAPLAN",
		Age:   24,
	}
	p2 := person{
		First: "Yasin",
		Last:  "Memiç",
		Age:   23,
	}
	people := []person{ p1, p2 }
	fmt.Println(people)

	// json.Marshal metodu hem JSON veriyi hem de hatayı döndürür.
	jsonVeri, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(jsonVeri)
	// Geri dönen JSON byte listesi biçiminde döneceği için eğer string 
	// 	olarak kullanmak istiyorsanız dönüştürmeniz gerekecektir.	
	fmt.Println(string(jsonVeri))
	
	

	// 003 JSON unmarshal
	s := `[{"first":"Kamil","last":"KAPLAN","age":24},{"first":"Yasin","last":"Memiç","age":23}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// peoplee := []person {}
	var peoplee []person
	err2 := json.Unmarshal(bs, &peoplee)
	if err2 != nil {
		fmt.Println(err2)	
	}
	fmt.Println("Tüm veri: ", peoplee)
	
	for  i, value :=range peoplee{
		fmt.Println("\nKişi Numarası", i)
		fmt.Println(value.First, value.Last, value.Age)
	}



	// 005 Sort
	xi := []int { 4,7,3,42,99,18,16,56,12}
	xs := []string {"Kamil", "Yasin", "Ali", "Melih", "Hasan", "Ayşe", "Rabia", "Duygu", "Dilan"}

	fmt.Println("Sort uygulanmadan önce:", xi)
	sort.Ints(xi)
	fmt.Println(xi)
	
	fmt.Println("-----")
	fmt.Println("Sort uygulanmadan önce:", xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println()
	sort.Strings(peoplee.First)
}	
