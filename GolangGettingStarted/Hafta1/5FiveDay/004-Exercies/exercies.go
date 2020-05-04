package main

import (
	"fmt"
	"encoding/json"
	"log"
)	
// 02 Exercies
type person struct {
	First     string
	Last      string
	Sayings []string
}

// 02 Exercies
type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func  foo(e error) {
	fmt.Println("foo ran -",e , "\n", e.(customErr).info)
}

// 03 Exercies
type sqrtError struct {
	lat string
	long string
	err error
}
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}
func  sqrt(f float64) (float64, error) {
	if f < 0 {
		// e := errors.New("more coffe needed")
		e := fmt.Errorf("more coffe neede - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}


func main() {
	// 01 Exercies
	p1 := person {
		First   : "Kamil",
		Last    : "KAPLAN",
		Sayings : []string {
			"Merhaba b1company.com.tr",
			"Merhaba gençlik",
			"Selam nasılsınız.",
		},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON marshal değil - Hata burada: ", err)
	}
	fmt.Println(string(bs))

	// 02 Exercies
	c1 := customErr{
		info : "need more coffe",
	}
	foo(c1)

	// 03 Exercies
	_, errS := sqrt(-10.23)
	if errS != nil {
		log.Println(errS)
	}
}

