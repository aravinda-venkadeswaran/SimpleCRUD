package service

import (
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/repository"
)

type PostService interface {
	Save(models.Post)
	SaveAll([]models.Post)
	Update(models.Post)
	FindAll() []models.Post
	FindByID(int) models.Post
	DeleteByID(int)
}

type PostServiceImpl struct {
}

func (service PostServiceImpl) Save(post models.Post) {
	repository.SaveToDB(&post)
}

func (service PostServiceImpl) SaveAll(posts []models.Post) {
	repository.SaveToDB(&posts)
}

func (service PostServiceImpl) Update(post models.Post) {
	var existingPost models.Post
	repository.FindByID(&existingPost, post.ID)
	if post.Title != "" {
		existingPost.Title = post.Title
	}
	if post.Body != "" {
		existingPost.Body = post.Body
	}
	repository.SaveToDB(&post)
}

func (service PostServiceImpl) DeleteByID(id int) {
	repository.DeleteFromDB(&models.Post{}, id)
}

func (service PostServiceImpl) FindAll() []models.Post {
	var posts []models.Post
	repository.FindAll(&posts, "Posts.Comments")
	return posts
}

func (service PostServiceImpl) FindByID(id int) models.Post {
	var post models.Post
	repository.FindByID(&post, id, "Posts.Comments")
	return post
}
