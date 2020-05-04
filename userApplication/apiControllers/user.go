package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/PROJECTS/user_application_go/dbConfig"
	"github.com/PROJECTS/user_application_go/libs"
	"github.com/PROJECTS/user_application_go/models"
	"net/http"
)

func GetUserList(c *gin.Context) {

	var users []models.Users
	err := dbConfig.DB.Find(&users).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserListById(c *gin.Context) {

	var user models.Users
	id := c.Param("id")

	err := dbConfig.DB.Where(id).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {

	var user models.Users
	//dbConfig.DB.CreateTable(user)
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "İşlem yaparken bir hata oluştu. Hata: "+err.Error())
		return
	}

	var dbUser models.Users
	dbConfig.DB.Where("email = ?", user.Email).First(&dbUser)

	if dbConfig.DB.NewRecord(&dbUser) {
		user.Hash = libs.GetMD5Hash(user.Hash)
		dbConfig.DB.Create(&user)
		c.JSON(http.StatusOK, "KAYIT İŞLEMİ BAŞARILI")
		return
	}

	c.JSON(http.StatusBadRequest, "KULLANICI KAYITLI")
}

func PutUser(c *gin.Context) {

	var user models.Users
	var dbUser models.Users

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "İşlem yaparken bir hata oluştu. Hata: "+err.Error())
		return
	}

	dbConfig.DB.Where("email = ?", user.Email).Find(&dbUser)

	if !dbConfig.DB.NewRecord(&dbUser) {
		dbUser.FirstName = user.FirstName
		dbUser.LastName = user.LastName
		dbUser.Phone = user.Phone
		dbUser.Age = user.Age
		dbUser.Company = user.Company
		dbUser.Active = user.Active
		dbUser.Address = user.Address
		if user.Hash != "" {
			dbUser.Hash = user.Hash
		}

		err := dbConfig.DB.Save(&dbUser).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, "İşlem yaparken bir hata oluştu. Hata: " + err.Error())
			return
		}

		c.JSON(http.StatusOK, "GÜNCELLEME İŞLEMİ BAŞARILI")
		return
	}

	c.JSON(http.StatusNotFound, "Böyle bir kullanıcı yok")

}

func DeleteUser(c *gin.Context) {
	var user models.Users
	id := c.Param("id")

	err := dbConfig.DB.Where(id).Delete(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: " + err.Error())
		return
	}
	
	c.JSON(http.StatusOK, "SİLME İŞLEMİ BAŞARILI")
}
