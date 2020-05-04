package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/PROJECTS/user_application_go/dbConfig"
	"github.com/PROJECTS/user_application_go/libs"
	"github.com/PROJECTS/user_application_go/models"
	"log"
	"net/http"
	"time"
)

func LoginUser(c *gin.Context) {
	var loginUser models.Users

	var form = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Hash     string `json:"-"` // hash diye bir value json'dan gelmiyor.
	}{}

	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, "İşlem yaparken bir hata oluştu. Hata: "+err.Error())
		return
	}

	form.Hash = libs.GetMD5Hash(form.Password)
	if form.Email == "" {
		c.JSON(http.StatusUnauthorized, "Kullanıcı adı girilmemiş")
		return
	}

	if form.Password == "" {
		c.JSON(http.StatusUnauthorized, "Şifre girilmemiş")
		return
	}

	dbConfig.DB.Where("email = ? and hash = ?", form.Email, form.Hash).Find(&loginUser)

	if loginUser.ID == 0 {
		c.JSON(http.StatusUnauthorized, "Kullanıcı adı veya şifre hatalı")
		return
	} else {
		strToken, err := libs.CreateJWT(loginUser)
		if err == nil {
			timeVal, _ := time.Parse("2006-01-02T15:04:05.000Z", libs.ParseToken(strToken, "expd").(string))
			c.JSON(http.StatusOK,
				models.LoginResp{
					TokenVal:  strToken,
					Expire:    timeVal,
					Email:     loginUser.Email,
				},
			)
			return
		} else {
			log.Println(err)
			c.JSON(http.StatusBadRequest, "Şifre oluşturma servisinde bir sorun var.")
		}
		return
	}
}
