package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPost(c *gin.Context) {

}
func createPost(c *gin.Context) {
	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request from user"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"data": newPost})
}
func updatePost(c *gin.Context) {

}
func main() {
	fmt.Println("Hello Go Gin")
	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello to Go Gin world"})
	})

	r.GET("/post",getPost)
	r.POST("/post",createPost)
	r.PUT("/post",updatePost)
	

	r.Run()
}