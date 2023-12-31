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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}
	defer file.Close()
	data, err := json.Marshal(newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error for writing data"})
		return
	}
	file.Write(data)
	c.JSON(http.StatusCreated,gin.H{"data": newPost})
}
func updatePost(c *gin.Context) {
	reqBody,_ := c.GetRawData()
	var data map[string]interface{}
	if err := json.Unmarshal(reqBody,&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Request data error"})
		return
	}
	fmt.Println("Request Data : ", data)
	if data["id"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Request ID is missing"})
		return
	}
	filename := fmt.Sprintf("./data/posts/%s_post.json",data["id"])
	fmt.Println("File Name : ", filename)
	file, err := os.OpenFile(filename, os.O_RDWR,0644)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Post id doesn't exist"})
		return
	}
	defer file.Close()
	fileData, _ := os.ReadFile(filename)
	var postData map[string]interface{}
	err = json.Unmarshal(fileData,&postData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Server error"})
		return
	}
	for key, value := range data {
		postData[key] = value
	}
	tempPost,_ := json.Marshal(postData)
	var updatedPost Post
	err = json.Unmarshal(tempPost,&updatedPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Unmarshal error"})
		return
	}
	byteData,_ := json.Marshal(updatedPost)
	file.Write(byteData)
	c.JSON(http.StatusAccepted, gin.H{"data" : updatedPost})
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
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Server error"})
		return
	}
	defer file.Close()
	data, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Server error for writing data"})
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