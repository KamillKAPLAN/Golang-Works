package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"io/ioutil"
	"log"
)	

func main() {
	// Dosya Yazma işlemi yapıyoruz.
	f, err := os.Create("names.txt")
	if err !=nil {
		fmt.Println(err)
		return
	}		
	defer f.Close()
		
	r := strings.NewReader("Melih KAPLAN")

	io.Copy(f, r)

	// Dosya okuma işlemi yapıyoruz.
	f2, err := os.Open("names.txt")
	if err !=nil {
		fmt.Println(err)
		return
	}		
	defer f2.Close()

	bs, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))

	// Dosya olup olmadığı kontrol
	_, err2 := os.Open("no-file.txt")
	if err2 != nil {
		// fmt.Println(err2)
		// log.Println(err2)
		log.Fatalln(err2)
		// log.Panic(err)
		// panic(err)
	}
}
