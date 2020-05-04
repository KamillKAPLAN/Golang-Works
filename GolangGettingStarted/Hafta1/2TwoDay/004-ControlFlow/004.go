package main

import "fmt"

func main() {
	// 001 For example Video1
	i := 1
	for i <= 3 {
		fmt.Print(" ", i)
		i = i + 1
	}
	fmt.Println()
	// 002 For example
	for j := 7; j <= 9; j++ {
		fmt.Print(" ", j)
	}
	fmt.Println()
	// 003 For example
	for {
		fmt.Println("loop")
		break
	}
	// 004 For example
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// 005 For example Video2
	for ii := 0; ii <= 9; ii++ {
		for jj := 0; jj <= 3; jj++ {
			fmt.Print(ii, jj)
		}
		fmt.Println()
	}
	// 006 For example Video3 and Video4
	ax := 1
	for {
		ax++
		if ax > 100 {
			break
		}
		if ax%2 != 0 {
			fmt.Println(ax)
			continue
		}
	}
	fmt.Println("done.")

	// 007 If example Video5
	if true && !false {
		fmt.Println("001")
	}
	if false && !true {
		fmt.Println("002")
	}
	// Önemli nir kullanım
	if dX := 42; dX == 2 {
		fmt.Println("003")
	}
	// 008 If Else example Video6
	x2 := 45
	if x2 == 40 {
		fmt.Println("değer: 40")
	} else if x2 == 42 {
		fmt.Println("değer: 42")
	} else if x2 == 44 {
		fmt.Println("değer: 44")
	} else if x2 == 45 {
		fmt.Println("değer: 45")
	} else {
		fmt.Println("değer 40, 42, 44 veya 45 değil")
	}

	// 009 For and If example Video 7
	for ij := 0; ij < 15; ij++ {
		if ij%2 == 0 {
			fmt.Println("2 ile bölümünen: ", ij)
		} else if ij%3 == 0 {
			fmt.Println("3 ile bölümünen: ", ij)
		} else {
			fmt.Println("Ne 2 nede 3 ile bölümünen: ", ij)
		}
	}

	// 010 Switch Statemet Video8
	n := "Kamil"
	switch n {
	case "Racep", "Kamil":
		fmt.Println("False sonuç")
	case "Ali":
		fmt.Println("2 == 4")
	case "Kamil":
		fmt.Println("4 == 4")
	default:
		fmt.Println("this is default")
	}
}
