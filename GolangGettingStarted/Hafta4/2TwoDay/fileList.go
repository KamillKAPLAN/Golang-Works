package main

import (
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("city/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == "code.data" {
			data, err := ioutil.ReadFile("code.data")
			if err != nil {
				log.Fatal(err)
			}
			log.Print(string(data))
		} else {
			mydata := []byte("All the data I wish to write to a file")
			err2 := ioutil.WriteFile("code.data", mydata, 0777)
			if err2 != nil {
				// print it out
				log.Fatal(err2)
			}
		}
	}
}
