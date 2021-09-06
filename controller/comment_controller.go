package controller

import (
	"fmt"
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service service.CommentService
}

func NewCommentController() *CommentController {
	commentService := service.CommentServiceImpl{}
	return &CommentController{service: commentService}
}

func (controller CommentController) GetAll(c *gin.Context) {
	fmt.Print(controller.service)
	comments := controller.service.FindAll()
	if len(comments) == 0 {
		c.Status(http.StatusNoContent)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Comments": comments,
		})
	}
}

func (controller CommentController) Get(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, _ := strconv.Atoi(idString)
	comment := controller.service.FindByID(id)
	c.JSON(http.StatusOK, comment)
}

func (controller CommentController) Add(c *gin.Context) {
	var comment models.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	controller.service.Save(comment)
	c.JSON(http.StatusOK, gin.H{
		"Comment": comment,
	})
}

func (controller CommentController) Update(c *gin.Context) {
	var comment models.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	controller.service.Save(comment)
	c.JSON(http.StatusOK, gin.H{
		"Comment": comment,
	})
}
