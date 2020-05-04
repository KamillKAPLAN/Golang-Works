package main

import (
	"fmt"
	"encoding/json"
	"os"
)

// 001 Example Marshal
type user struct {
	First string
	Age   int
}

// 002 Example Unmarshal
type person struct {
	First   string    `json:"First"`
	Last    string    `json:"Last"`
	Age     int       `json:"Age"`
	Job     string    `json:"Job"`
}

func main() {

	// 001 Example Marshal
	u1 := user {
		First : "Kamil",
		Age   : 24,
	}
	u2 := user {
		First : "Melih",
		Age   : 23,
	}
	u3 := user {
		First : "Yasin",
		Age   : 23,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	value, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(value))

	fmt.Println("-----------------")

	// 002 Example Unmarshal
	//p1 :=person {
	//	First : "Kamil",
	//	Last  : "KAPLAN",
	//	Age   : 24,
	//	Job   : "Yazılım Mühendisi",
	//}
	//p2 :=person {
	//	First : "İbrahim",
	//	Last  : "",
	//	Age   :  42,
	//	Job   : "Patron",
	//}
	//p3 :=person {
	//	First : "Melih",
	//	Last  : "KAPLAN",
	//	Age   :  23,
	//	Job   : "Elektrik Elektronik Tekniker",
	//}

	// people := []person { p1, p2, p3 }

	s:=`[{"First":"Kamil","Last":"KAPLAN","Age":24,"Job":"Yazılım Mühendisi"},{"First":"İbrahim","Last":"","Age":42,"Job":"Patron"},{"First":"Melih","Last":"KAPLAN","Age":23,"Job":"Elektrik Elektronik Tekniker"}]`
	fmt.Println(s)	

	var people []person

	err2 := json.Unmarshal([]byte(s),&people)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", v.First, v.Last, v.Age, v.Job)
	}

	//003 Example Encoder

	err3 := json.NewEncoder(os.Stdout).Encode(people)
	if err3 != nil {
		fmt.Println(err3)
	}

	
}