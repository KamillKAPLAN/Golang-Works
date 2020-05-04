package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html") // veya "templates/view/*.html"

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")

		timesOff, ok := TimesOff[id]

		 if !ok {
			c.String(http.StatusNotFound,"404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overvide.html",
			map[string]interface{}{ // map[string]interface{} bunun yerine gin.H 'de yazılabilir.
				"TimesOff" : timesOff,

			})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin" : "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"admin-overview.html",nil)
	})

	// r.Static("/template","./template")

	return r
}