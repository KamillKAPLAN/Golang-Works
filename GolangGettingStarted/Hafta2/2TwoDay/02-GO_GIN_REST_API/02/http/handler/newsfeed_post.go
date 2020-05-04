package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/GO/Hafta2/2TwoDay/02-GO_GIN_REST_API/02/platform/newsfeed"
	"net/http"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
} 

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc{
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title:requestBody.Title,
			Post:requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}