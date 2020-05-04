package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*
	günlük dosya oluşturuyor. ve var olan dosyalar içinde geçmiş tarihli olanı siliyor.
*/

func main() {
	now := time.Now().Add(time.Hour * 24)
	dateString := now.Format("2006-01-02")

	beforeDay := now.AddDate(0, 0, -1)
	beforeDateString := beforeDay.Format("2006-01-02")
	isFile := false

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	fileName := dateString + ".json"
	for _, f := range files {
		if f.Name() == fileName {
			data, err := ioutil.ReadFile("./" + fileName)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(data))
			os.Remove((beforeDateString) + ".forecast" + ".json")
			isFile = true
			break
		}
	}

	if !isFile {
		err = ioutil.WriteFile((dateString)+".forecast"+".json", []byte("Hello"), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
}
