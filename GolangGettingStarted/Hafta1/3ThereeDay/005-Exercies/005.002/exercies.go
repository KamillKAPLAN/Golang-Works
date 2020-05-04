package main

import (
	"fmt"
	"math"
)

type daire struct {
	radius float64
}

type kare struct {
	length float64
}

func (c daire) alan() float64 {
	fmt.Print("Dairenin Alanı: ")
	return math.Pi * c.radius * c.radius
}
func (s kare) alan() float64 {
	fmt.Print("Karenin Alanı: ")
	return s.length * s.length
}

type sekil interface {
	alan() float64
}

func bilgi(s sekil) {
	fmt.Println(s.alan())
}

func main() {
	daire := daire{
		radius: 12.345,
	}
	kare := kare{
		length: 15,
	}
	bilgi(daire)
	bilgi(kare)
}
