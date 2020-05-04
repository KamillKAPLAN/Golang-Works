package main

import (
	"fmt"
	"github.com/Employee.API/apiRoots"
	"github.com/Employee.API/dbConfig"
	"github.com/Employee.API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	dbConfig.InitDB()

	if !dbConfig.DB.HasTable(models.Employee{}) {
		dbConfig.DB.CreateTable(models.Employee{})
		fmt.Println("Employee table created")
	}

	api := app.Group("/api")

	apiRoots.EmployeeAPI(api)

	_ = app.Run(fmt.Sprintf(":1111"))
}
