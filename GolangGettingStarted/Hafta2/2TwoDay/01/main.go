package main

import (
	"github.com/GO/Hafta2/2TwoDay/01/apiControllers"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	api :=app.Group("/api",)
	apiControllers.BabanGET(api)
	apiControllers.AnanDELETE(api)
	apiControllers.AnanPOST(api)
	app.Run()

}