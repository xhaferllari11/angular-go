package main

import (
	"github.com/gin-gonic/gin"
	"workingWithAPIs/httpd/handlers"
	"workingWithAPIs/platform/newsfeed"
)

func main() {
	r := gin.Default()
	feed := newsfeed.New()

	r.GET("/ping", handlers.PingGet)
	r.GET("/newsfeed", handlers.NewsfeedGet(feed))
	r.POST("/newsfeed", handlers.NewsfeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//feed := newsfeed.New()
	//fmt.Println(feed)
	//feed.Add(newsfeed.Item{Title: "Alban", Post: "is Bosss"})
	//fmt.Println(feed)
}
