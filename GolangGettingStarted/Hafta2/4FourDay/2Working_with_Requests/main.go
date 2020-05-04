package main

import (
	"log"
	"strings"
)

func main() {
	var companyName = "B1 Yönetim Sislemleri Yazılım ve Danışmanlık Ltd. Şti"

	companyStrArray := strings.Split(companyName, " ")

	var companyBuff string
	var name1 string
	var name2 string
	for _, item := range companyStrArray {

		if companyBuff != "" {
			companyBuff += " "
		}

		companyBuff += item

		if len(companyBuff) <= 40 {
			name1 = companyBuff
			//log.Println("**",companyBuff,len(companyBuff))
		} else {
			//log.Println("hooppa")
			name2 = companyName[len(name1)+1:]
			log.Println(name1)
			log.Println(name2)
			break
		}

		//log.Println("Name1: ", name1)

	}
}
