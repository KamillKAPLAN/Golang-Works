package baba

import (
	"github.com/GO/Hafta2/1OneDay/03/dede"
	"github.com/gin-gonic/gin"
)

func Baba(api *gin.RouterGroup) {
	api.GET("/baba:istanbul", dede.Dede)
}
