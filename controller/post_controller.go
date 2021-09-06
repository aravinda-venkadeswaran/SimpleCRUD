package controller

import (
	"fmt"
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service service.PostService
}

func NewPostController() *PostController {
	postService := service.PostServiceImpl{}
	return &PostController{service: postService}
}

func (controller PostController) GetAll(c *gin.Context) {
	fmt.Print(controller.service)
	posts := controller.service.FindAll()
	if len(posts) == 0 {
		c.Status(http.StatusNoContent)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Posts": posts,
		})
	}
}

func (controller PostController) Get(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, _ := strconv.Atoi(idString)
	post := controller.service.FindByID(id)
	c.JSON(http.StatusOK, post)
}

func (controller PostController) Add(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	controller.service.Save(post)
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}

func (controller PostController) Update(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	controller.service.Save(post)
	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}
