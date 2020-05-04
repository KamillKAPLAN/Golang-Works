package apiControllers

import (
	"github.com/GO/Hafta2/2TwoDay/01/apiRoots"
	"github.com/gin-gonic/gin"
)

func BabanGET(api *gin.RouterGroup) {
	api.GET("/babanGET", apiRoots.DedenGET)

}

func AnanDELETE(api *gin.RouterGroup) {
	api.DELETE("/ananDELETE", apiRoots.NinenDELETE)
}

func AnanPOST(api *gin.RouterGroup){
	api.POST("/ananPOST",apiRoots.NinenPOST)
}