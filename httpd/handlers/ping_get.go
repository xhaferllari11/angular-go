package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "found me",
	})
}