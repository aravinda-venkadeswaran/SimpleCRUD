package main

import (
	"encoding/json"
	"io/ioutil"
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/repository"
	"localSandbox/SimpleCRUD/routes"
	"localSandbox/SimpleCRUD/service"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB

func fetchUsers() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)
	var users []models.User
	json.Unmarshal([]byte(stringBody), &users)
	var userService service.UserService = service.UserServiceImpl{}
	userService.SaveAll(users)
}

func fetchPosts() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)
	var posts []models.Post
	json.Unmarshal([]byte(stringBody), &posts)
	var postService service.PostService = service.PostServiceImpl{}
	postService.SaveAll(posts)
}

func fetchComments() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)
	var comments []models.Comment
	json.Unmarshal([]byte(stringBody), &comments)
	var commentService service.CommentService = service.CommentServiceImpl{}
	commentService.SaveAll(comments)
}

func main() {

	repository.InitializeDatabase()
	fetchUsers()
	fetchPosts()
	fetchComments()
	routes.InitializeRoutes()

}
