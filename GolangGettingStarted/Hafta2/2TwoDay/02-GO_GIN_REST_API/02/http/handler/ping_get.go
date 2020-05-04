package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func PingGet(c *gin.Context) { --- 1
//	c.JSON(http.StatusOK, map[string]string{
//		"hello" : "Found me",
//	})
//}

func PingGet() gin.HandlerFunc{ // --- 2
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello" : "Found me",
		})
	}
}