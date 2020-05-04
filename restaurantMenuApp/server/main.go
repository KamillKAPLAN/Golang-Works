package main

import (
	"github.com/ridvansumset/restaurant-menu-app/server/controllers"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()
	controllers.Init()

	router.Pre(echomiddleware.RemoveTrailingSlash())
	router.Use(echomiddleware.Logger())
	router.Use(echomiddleware.Recover())
	router.Use(echomiddleware.Gzip())

	router.Use(
		echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
			AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
			AllowOrigins:     []string{"*"},
			AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}))

	router.GET("/categories", controllers.ListCategories)
	router.GET("/categories/:category_id", controllers.GetCategory)
	router.POST("/categories", controllers.CreateCategory)
	router.DELETE("/categories/:category_id", controllers.DeleteCategory)
	router.PUT("/categories/:category_id", controllers.UpdateCategory)

	router.GET("/categories/:category_id/products", controllers.ListProducts)
	router.GET("/categories/:category_id/products/:product_id", controllers.GetProduct)
	router.POST("/categories/:category_id/products", controllers.CreateProduct)
	// router.DELETE("/categories/:category_id/products/:product_id", controllers.DeleteProduct)
	// router.PUT("/categories/:category_id/products/:product_id", controllers.UpdateProduct)

	router.GET("/categories/:category_id/products/:product_id/options", controllers.ListOptions)
	router.GET("/categories/:category_id/products/:product_id/options/:option_id", controllers.GetOption)
	// router.POST("/categories/:category_id/products", controllers.CreateOption)
	// router.DELETE("/categories/:category_id/products/:product_id", controllers.DeleteOption)
	// router.PUT("/categories/:category_id/products/:product_id", controllers.UpdateOption)

	router.GET("/categories/:category_id/products/:product_id/options/:option_id/choices", controllers.ListChoices)
	router.GET("/categories/:category_id/products/:product_id/options/:option_id/choices/:choice_id", controllers.GetChoice)
	// router.POST("/categories/:category_id/products", controllers.CreateOption)
	// router.DELETE("/categories/:category_id/products/:product_id", controllers.DeleteOption)
	// router.PUT("/categories/:category_id/products/:product_id", controllers.UpdateOption)

	router.Logger.Fatal(router.Start(":1323"))
}
