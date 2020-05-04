package main

import "fmt"

func main() {
	// for exercies

	// 001
	for i := 0; i <= 100; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println()
	// 002
	for j := 65; j <= 90; j++ {
		fmt.Printf("%#U\n", j)
	}
	// 003
	a := 1994
	for a <= 2019 {
		fmt.Print(" - ", a)
		a++
	}
	// 004
	b := 1985
	for {
		if b > 1999 {
			break
		}
		fmt.Println(b)
		b++
	}

	// switch exercies

	// 001
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains")
	case "swimming":
		fmt.Println("go to the pool")
	case "surfing":
		fmt.Println("go to the hawaii")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
