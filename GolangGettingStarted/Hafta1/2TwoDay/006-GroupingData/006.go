package main

import "fmt"

func main() {
	// 001 Array
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println("x lenght: ", len(x))

	// 002 Slice -- slice aynı türdeki değerleri birlikte gruplamanıza olanak verir
	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y) // y := []type {values} // birleşik hazır bilgi

	// Slice - for range
	for index, value := range y {
		fmt.Println(index, value)
	}

	// Slice - slicing a slice
	y1 := []int{4, 5, 7, 8, 42}
	fmt.Println(y1[1])
	fmt.Println(y1)
	fmt.Println(y1[1:])  // 1. indisden itibaren göster
	fmt.Println(y1[1:3]) // 1.indis dahil 3. indis dahil değil
	for i, v := range y1 {
		fmt.Println(i, v)
	}
	//for j := 0; j <= len(y1); j++ {
	//	fmt.Println(j, y1[j])
	//}

	// Slice - append to a slice
	y1 = append(y, 6, 12, 95, 65, 56)
	fmt.Println(y1)
	// bir slice ı diğer bir slice'a eklemek istersek aşağıdaki kullanumı yaparız.
	y = append(y, y1...)
	fmt.Println(y)

	// Slice - deleting from a slice
	fmt.Println(y[:2]) // 2.indise kadar al
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	// Slice - multi-dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)

	// 003 Map - introduction
	m := map[string]int{
		"Kamil":   25,
		"learnGO": 24,
	}
	fmt.Println(m)
	fmt.Println(m["Kamil"])
	fmt.Println(m["x"])
	v, ok := m["x"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["learnGO"]; ok {
		fmt.Println("This is the if print", v)
	}

	// Map - add element & range
	m["tood"] = 33
	for a, b := range m {
		fmt.Println(a, b)
	}

	// Map - delete
	delete(m, "Kamil")
	fmt.Println(m)
	if v, ok := m["learnGO"]; ok {
		fmt.Println("value: ", v)
		delete(m, "learnGO")
	}
	fmt.Println(m)
}
