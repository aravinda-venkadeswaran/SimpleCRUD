package service

import (
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/repository"
)

type UserService interface {
	Save(models.User)
	SaveAll([]models.User)
	Update(models.User)
	FindAll() []models.User
	FindByID(int) models.User
	DeleteByID(int)
}

type UserServiceImpl struct {
}

func (service UserServiceImpl) Save(user models.User) {
	repository.SaveToDB(&user)
}

func (service UserServiceImpl) SaveAll(users []models.User) {
	repository.SaveToDB(&users)
}

func (service UserServiceImpl) Update(user models.User) {
	var existingUser models.User
	repository.FindByID(&existingUser, user.ID)
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Address != (models.Address{}) {
		existingUser.Address = user.Address
	}
	if user.Phone != "" {
		existingUser.Phone = user.Phone
	}
	if user.Website != "" {
		existingUser.Website = user.Website
	}
	if user.Company != (models.Company{}) {
		existingUser.Company = user.Company
	}
	repository.SaveToDB(&user)
}

func (service UserServiceImpl) DeleteByID(id int) {
	repository.DeleteFromDB(&models.User{}, id)
}

func (service UserServiceImpl) FindAll() []models.User {
	var users []models.User
	repository.FindAll(&users, "Posts.Comments")
	return users
}

func (service UserServiceImpl) FindByID(id int) models.User {
	var user models.User
	repository.FindByID(&user, id, "Posts.Comments")
	return user
}
