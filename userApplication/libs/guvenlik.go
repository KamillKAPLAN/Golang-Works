package libs

import (
	"crypto/md5"
	"encoding/hex"
	dbconfig "github.com/PROJECTS/user_application_go/dbConfig"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func BiletKontrol(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		//c.AbortWithStatus(http.StatusOK)
	}

	var tokenStr string = c.Request.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "bearer ", "", 10)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(dbconfig.MySecretKey), nil
	})
	if err == nil && token.Valid {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
		c.Abort()
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
