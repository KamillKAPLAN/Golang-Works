package main

import "fmt"

func main() {
	n, err := fmt.Println("Kamil KAPLAN", 42, true)
	fmt.Println(n, " ", err)
	m, _ := fmt.Println("Kamil KAPLAN", 42, true)
	fmt.Println(m)
}
