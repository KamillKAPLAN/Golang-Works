package apiControllers

import (
	"github.com/Employee.API/dbConfig"
	"github.com/Employee.API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostEmployee(c *gin.Context) {

}

func GetEmployee(c *gin.Context) {
	var employees []models.Employee
	err := dbConfig.DB.Find(&employees).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, &employees)
}

func GetEmployeeById(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	err := dbConfig.DB.Where(id).Find(&employee).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, &employee)
}

func PutEmployee(c *gin.Context) {

}

func DeleteEmployees(c *gin.Context) {
	var employee models.Employee

	err := dbConfig.DB.Delete(&employee).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "Silme İşlemi Başarılı")
}

func DeleteEmployeeById(c *gin.Context) {
	var employee models.Employee

	err := dbConfig.DB.Where(id).Delete(&employee).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "Silme İşlemi Başarılı")

}