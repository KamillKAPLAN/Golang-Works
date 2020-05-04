package main

import (
	"fmt"
	"strings"
)

func main() {
	val := "zeytin"
	val = strings.Replace(val, "a", "A", 1)
	val = strings.Replace(val, "b", "B", 1)
	val = strings.Replace(val, "c", "C", 1)
	val = strings.Replace(val, "ç", "Ç", 1)
	val = strings.Replace(val, "d", "D", 1)
	val = strings.Replace(val, "e", "E", 1)
	val = strings.Replace(val, "f", "F", 1)
	val = strings.Replace(val, "g", "G", 1)
	val = strings.Replace(val, "ğ", "Ğ", 1)
	val = strings.Replace(val, "h", "H", 1)
	val = strings.Replace(val, "ı", "I", 1)
	val = strings.Replace(val, "i", "İ", 1)
	val = strings.Replace(val, "j", "J", 1)
	val = strings.Replace(val, "k", "K", 1)
	val = strings.Replace(val, "l", "L", 1)
	val = strings.Replace(val, "m", "M", 1)
	val = strings.Replace(val, "n", "N", 1)
	val = strings.Replace(val, "o", "O", 1)
	val = strings.Replace(val, "ö", "Ö", 1)
	val = strings.Replace(val, "p", "P", 1)
	val = strings.Replace(val, "r", "R", 1)
	val = strings.Replace(val, "s", "S", 1)
	val = strings.Replace(val, "ş", "Ş", 1)
	val = strings.Replace(val, "t", "T", 1)
	val = strings.Replace(val, "u", "U", 1)
	val = strings.Replace(val, "ü", "Ü", 1)
	val = strings.Replace(val, "v", "V", 1)
	val = strings.Replace(val, "y", "Y", 1)
	val = strings.Replace(val, "z", "Z", 1)
	fmt.Println(val,"\n","------------------------")
}
