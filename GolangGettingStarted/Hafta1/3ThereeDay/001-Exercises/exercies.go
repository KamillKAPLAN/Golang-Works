package main

import "fmt"

func main() {
	// 001 Example
	x := [5]int{12, 5, 8, 14, 6}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	// 002 Example
	x2 := []int{12, 5, 8, 14, 6, 7, 9, 1, 0, 14, 25, 36}
	fmt.Println(x2)

	// 003 Example
	x3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(x3[:5])  // ilk 5 indis'i göster
	fmt.Println(x3[5:])  // ilk 5 indis'den sonrasını göster
	fmt.Println(x3[2:7]) // 2. indis ile 7. indis arasını göster
	fmt.Println(x3[1:6])
	fmt.Println(x3)

	// 004 Example
	x4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x4)
	x4 = append(x4, 52)
	fmt.Println(x4)
	x4 = append(x4, 53, 54, 55)
	fmt.Println(x4)

	y4 := []int{56, 57, 58, 59, 60}
	x4 = append(x4, y4...)
	fmt.Println(x4)

	// 005 Example
	x5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x5)
	x5 = append(x5[:3], x5[6:]...)
	fmt.Println(x5)

	// 006 Example
	xs1 := []string{"Kamil", "KAPLAN", "Elazığ", "Fırat", "Üniversitesi"}
	xs2 := []string{"Teknoloji", "Fakültesi", "Yazılım", "Mühendisliği", "Yeni MEZUN"}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Println("kayıt: ", i)
		for j, k := range v {
			fmt.Printf("\t index pozisyonu: %v \t değer: \t %v \n", j, k)
		}
	}

	// 007 Example
	m := map[string][]string{
		"a": []string{"a1", "a2", "a3"},
		"b": []string{"b1", "b2", "b3"},
		"c": []string{"c1", "c2", "c3"},
	}
	fmt.Println(m)
	for i, v := range m {
		fmt.Println("Bu kayıt: ", i)
		for i2, v2 := range v {
			fmt.Println("\t", i2, v2)
		}
	}

	// 008 Example
	m["d"] = []string{"d1", "d2", "d3"} // d key ve value değerleri m  map'ine eklendi

	// 009 Example
	delete(m, "a") // m map'inden a key'ini sildik
	fmt.Println(m)
}
