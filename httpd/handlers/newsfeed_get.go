package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"workingWithAPIs/platform/newsfeed"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{

	return func(c *gin.Context) {
		fmt.Println(feed)
		result := feed.GetAll()
		fmt.Println(result)
		c.JSON(http.StatusOK, result)
	}
}