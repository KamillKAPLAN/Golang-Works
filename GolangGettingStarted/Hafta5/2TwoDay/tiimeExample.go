package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Geçerli Zaman ve Saat : ", currentTime.String())
	fmt.Println("Geçerli Zaman Format (YYYY-MM-DD) : ", currentTime.Format("2006-01-02"))
	fmt.Println("Geçerli Zaman ve Saat Format (YYYY-MM-DD hh:mm:ss) :  ", currentTime.Format("2006-01-02T15:04:05"))
	format := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(format)
	exampleDate := "2019-07-30T000:00+03:00"
	nowDate := currentTime.Format("2006-01-02T15:04:05")

	if exampleDate < nowDate {
		log.Println("Sistem küçük")
	} else {
		log.Println("Bam bam bam")
	}

	// only time.Clock()
	format2 := currentTime.Format("2006-01-02T15:04:05")
	log.Println("asddasdasdasd: ", format2)
	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, format2)
	if err != nil {
		log.Println(err, " : bam bam bam bam ")
	}
	hr, _, _ := timeStamp.Clock()
	log.Println("time Clock : ", hr)

	aformat := currentTime.Format("15:04:05")
	log.Println(aformat)

	beginTime := "07:00:00"
	endTime := "19:00:00"
	if beginTime < aformat && aformat < endTime {
		log.Println("özellik tamam")
	} else {
		log.Println("ne yapıon")
	}

	hour, d, d := time.Now().Clock()

	log.Println("saaaatttt",ahour)

}
