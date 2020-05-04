package dede

import (
	"net/http"

	"github.com/GO/Hafta2/1OneDay/03/model"

	"github.com/gin-gonic/gin"
)

var atam model.Ata

func Dede(c *gin.Context) {
	atam.Sehir = c.Params.ByName("istanbul")
	atam.Isim = "kamil"
	atam.Message = "Atasanı buldu"
	var i int
	for i = 0; i < len(atam.Sehir); i++ {
		i = i + 1
	}
	atam.Sayi = i
	c.JSON(http.StatusOK, atam)
}
