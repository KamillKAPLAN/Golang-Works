package apiRoots

import (
	"github.com/GO/Hafta2/2TwoDay/01/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var atam models.BuyukEbeveyn
func DedenGET(c *gin.Context) {
	atam.FName = "Ahmet"
	atam.LName = "KAPLAN"
	atam.Age = 75
	atam.Messagge = "Atasını buldu"
	atam.Sehir = "Malatya"
	atam.AtamCokYasa = len(atam.FName)
	var tplm int = 0
	for i := 0; i <= atam.AtamCokYasa; i++ {
		tplm = tplm + i
	}
	atam.AtamCokYasa = tplm
	c.JSON(http.StatusOK, atam)
}

var atamGorev [6]models.BuyukEbeveyn
func NinenDELETE(c *gin.Context) {
	for i:=range atamGorev{
		atamGorev[i].Gorev = "SİLİNDİ"
	}
	c.JSON(http.StatusOK, atamGorev)
}

func NinenPOST(c *gin.Context){
	fName := []string{"Ahmet", "Ali", "Kamil", "Yasin", "Serkan", "Hüseyin"}
	lName := []string{"Koç", "Aslan", "KAPLAN", "MEMİÇ", "Doğangüneş", "Bilmiyorum"}
	age := []int{43, 42, 25, 24, 38, 24}
	message := []string{"hi", "how are you", "Selam Gençlik", "Hoop", "Kardeşim borcunu öde", "Ayıp ettin herşeyi hallederiz."}
	sehir := []string{"Malatya", "Elazığ", "Malatya", "Giresun", "Malatya", "Malatya"}
	gorev := []string{"Emekli", "Çalışan", "Mühendis", "Mühendis", "İşçi", "Mühendis"}
	for i:=range atamGorev {
		atamGorev[i].FName = fName[i]
		atamGorev[i].LName = lName[i]
		atamGorev[i].Age= age[i]
		atamGorev[i].Messagge = message[i]
		atamGorev[i].Sehir = sehir[i]
		atamGorev[i].AtamCokYasa = len(fName[i])
		atamGorev[i].Gorev = gorev[i]
	}
	c.JSON(http.StatusOK, atamGorev)
}


