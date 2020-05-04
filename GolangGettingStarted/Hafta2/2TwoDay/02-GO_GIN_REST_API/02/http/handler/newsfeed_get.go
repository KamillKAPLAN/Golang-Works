package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/GO/Hafta2/2TwoDay/02-GO_GIN_REST_API/02/platform/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}