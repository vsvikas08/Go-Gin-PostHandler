package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello Go Gin")
	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello to Go Gin world"})
	})
	r.Run()
}