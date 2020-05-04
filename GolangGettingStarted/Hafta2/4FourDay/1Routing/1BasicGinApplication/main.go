package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Kamil KAPLAN")
	})

	r.Run(":3000")
}
