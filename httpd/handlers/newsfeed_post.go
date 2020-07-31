package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"workingWithAPIs/platform/newsfeed"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc{
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		fmt.Println("testing when this printed")
		c.Bind(&requestBody)
		fmt.Println(c)
		fmt.Println(requestBody)
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}
		fmt.Println(item)
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}