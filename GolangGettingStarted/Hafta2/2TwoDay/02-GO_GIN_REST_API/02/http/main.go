package main

import (
	"github.com/GO/Hafta2/2TwoDay/02-GO_GIN_REST_API/02/http/handler"
	"github.com/GO/Hafta2/2TwoDay/02-GO_GIN_REST_API/02/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {

	feed := newsfeed.New()
    r := gin.Default()
	// r.GET("/ping", handler.PingGet) 1

	r.GET("/ping", handler.PingGet()) // 2
	r.GET("/newsfeedGET", handler.NewsfeedGet(feed))
	r.POST("/newsfeedPOST", handler.NewsfeedPost(feed))

	r.Run()

//	feed := newsfeed.New()
//	fmt.Println(feed)
//	feed.Add(newsfeed.Item{"Hello","what do you do?"})
//	fmt.Println(feed)
}