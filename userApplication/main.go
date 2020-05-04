package main

import (
	"fmt"
	"github.com/PROJECTS/user_application_go/apiControllers"
	"github.com/PROJECTS/user_application_go/apiRoots"
	"github.com/PROJECTS/user_application_go/dbConfig"
	"github.com/PROJECTS/user_application_go/libs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	app := gin.Default()

	app.Use(cors.New(getDefaultCorsConfig()))

	app.POST("/login", apiControllers.LoginUser)

	app.POST("/user", apiControllers.PostUser)

	dbConfig.InitDB()

	api := app.Group("/api", libs.BiletKontrol)

	apiRoots.UserApi(api)

	_ = app.Run(fmt.Sprintf(":1111"))
}

func getDefaultCorsConfig() cors.Config {

	return cors.Config{
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Bilet"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	}

}
