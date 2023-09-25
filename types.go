package main

type Post struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Date     string    `json:"date"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	PostId  string `json:"post_id"`
	User    User   `json:"user"`
	Comment string `json:"comment"`
}

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string `json:"password"`
	Image    string `json:"img,omitempty"`
}