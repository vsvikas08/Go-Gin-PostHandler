package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getPost(c *gin.Context) {

}
func createPost(c *gin.Context) {
	var newPost Post
	reqBody,_ := c.GetRawData()
	if err := json.Unmarshal(reqBody,&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request from user"})
		return
	}
	t := time.Now()
	id := strconv.FormatInt(t.Unix(),10)
	newPost.Id = id
	newPost.Date = t.Format("2006-01-02")

	fileName := fmt.Sprintf("./data/posts/%s_post.json",id)
	file, err := os.Create(fileName)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Server error"})
		return
	}
	defer file.Close()
	data, err := json.Marshal(newPost)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "server error for writing data"})
		return
	}
	file.Write(data)
	c.JSON(http.StatusCreated,gin.H{"data": newPost})
}
func updatePost(c *gin.Context) {

}
func getUserById(c *gin.Context) {}
func createNewUser(c *gin.Context) {
	var newUser User
	reqBody,_ := c.GetRawData()
	if err := json.Unmarshal(reqBody,&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body"})
		return
	}
	if newUser.Email == "" || newUser.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name, Email  or Password is missing."})
		return
	} 
	id := strconv.FormatInt(time.Now().Unix(),10)
	newUser.Id = string(id)
	fileName := fmt.Sprintf("./data/users/%s_user.json",id)
	file, err :=  os.Create(fileName)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message" : "Server error"})
		return
	}
	defer file.Close()
	data, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message" : "Server error for writing data"})
		return
	}
	file.Write(data)
	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}
func commentOnPost(c *gin.Context) {
	
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

	r.GET("/user/:id",getUserById)
	r.POST("/create_user",createNewUser)

	r.POST("/comment/:pst_id",commentOnPost)
	

	r.Run()
}