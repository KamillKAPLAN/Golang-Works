package main

import (
	"fmt"

	"github.com/GO/Hafta2/1OneDay/03/ana"

	"github.com/GO/Hafta2/1OneDay/03/baba"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	api := app.Group("/api")
	baba.Baba(api)
	ana.Ana(api)
	app.Run(fmt.Sprintf(":%d", 1456))
}
