package apiRoots

import (
	"github.com/Employee.API/apiControllers"
	"github.com/gin-gonic/gin"
)

func EmployeeAPI(api *gin.RouterGroup) {
	// CreateReadUpdateDelete = CRUD

	// Create employee model                     (CREATE)
	api.POST("/employeePOST", apiControllers.PostEmployee)       // localhost:1111/api//employeePOST
	// Read employee model                       (READ)
	api.GET("/employeesGET", apiControllers.GetEmployee)         // localhost:1111/api//employeeGET
	// Read id option employee model             (READ)
	api.GET("/employeeGET/:id", apiControllers.GetEmployeeById)    // localhost:1111/api//employeeGET:id
	// Update employee model                     (UPDATE)
	api.PUT("/employeeUpdate", apiControllers.PutEmployee)       // localhost:1111/api//employeeUpdate
	// Delete employee model                     (DELETE)
	api.DELETE("/employeesDELETE", apiControllers.DeleteEmployees) // localhost:1111/api//employeeDELETE
	// Delete id option employee model           (DELETE)
	api.DELETE("/employeeDELETE:id", apiControllers.DeleteEmployeeById) // localhost:1111/api//employeeDELETE
}
