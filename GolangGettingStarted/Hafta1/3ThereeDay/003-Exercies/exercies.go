package main

import "fmt"

type person struct {
	first    string
	last     string
	favTatli []string
}

// 003 Example struct
type arac struct {
	kapi int
	renk string
}
type kamyon struct {
	arac
	dortTeker bool
}
type sedan struct {
	arac
	luks bool
}

func main() {
	// 001 Example
	p1 := person{
		first: "Kamil",
		last:  "KAPLAN",
		favTatli: []string{
			"a",
			"b",
			"c",
		},
	}
	fmt.Println(p1)
	fmt.Printf("Name: %v \t LastName: %v \n", p1.first, p1.last)
	for i, v := range p1.favTatli {
		fmt.Println(i, v)

	}

	// 002 Example
	m := map[string]person{
		p1.last: p1,
	}
	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favTatli {
			fmt.Println(i, val)
		}
	}

	// 003 Example
	k := kamyon{
		arac: arac{
			kapi: 2,
			renk: "beyaz",
		},
		dortTeker: true,
	}
	s := sedan{
		arac: arac{
			kapi: 4,
			renk: "gümüş",
		},
		luks: false,
	}
	fmt.Println(k)
	fmt.Println(s)

	// 004 Example
	strct := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "Kamil",
		friends: map[string]int{
			"a": 5,
			"b": 6,
			"c": 7,
		},
		favDrink: []string{
			"su",
			"ayran",
		},
	}
	for i, v := range strct.favDrink {
		fmt.Println(i, v)
	}
	for _, val := range strct.friends {
		fmt.Println(val)
	}

}
