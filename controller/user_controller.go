package controller

import (
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAll(c *gin.Context)
	// Save(ctx *gin.Context) error
	// Update(ctx *gin.Context) error
	// Delete(ctx *gin.Context) error
	// ShowAll(ctx *gin.Context)
}

type UserController struct {
	service service.UserService
}

func NewUserController() *UserController {
	userService := service.UserServiceImpl{}
	return &UserController{service: userService}
}

func (controller UserController) GetAll(c *gin.Context) {
	users := controller.service.FindAll()
	if len(users) == 0 {
		c.Status(http.StatusNoContent)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Users": users,
		})
	}
}

func (controller UserController) Get(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, _ := strconv.Atoi(idString)
	user := controller.service.FindByID(id)
	c.JSON(http.StatusOK, user)
}

func (controller UserController) Add(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	controller.service.Save(user)
	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func (controller UserController) Update(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	controller.service.Save(user)
	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func (controller UserController) Delete(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, _ := strconv.Atoi(idString)
	controller.service.DeleteByID(id)
	c.Status(http.StatusOK)
}
