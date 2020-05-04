package _1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GETHomePage(context *gin.Context) {
	context.JSON(200, gin.H{
		"message" : "GET Home Page",
	})
}

func POSTHomePage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message" : "POST Home Page",
	})
}

func QueryStrings(ctx *gin.Context) {
	name := ctx.Query("name")
	age  := ctx.Query("age")

	ctx.JSON(http.StatusOK, gin.H{
		"name" : name,
		"age"  : age,
	})
}

func PathParameters(ctx *gin.Context) {
	name := ctx.Param("name")
	age  := ctx.Param("age")

	ctx.JSON(http.StatusOK, gin.H{
		"name" : name,
		"age"  : age,
	})
}

func POSTBody(ctx *gin.Context) {
	body := ctx.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message" : string(value),
	})
}

func main() {
	app := gin.Default()
	app.GET("/", GETHomePage)
	app.POST("/",POSTHomePage)
	app.GET("/query", QueryStrings)             // query?name=kamil&age=25
	app.GET("/path/:name/:age", PathParameters) // path/kamil/25
	app.POST("/body", POSTBody)
	app.Run(fmt.Sprint(":1111"))
}

