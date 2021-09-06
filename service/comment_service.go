package service

import (
	"localSandbox/SimpleCRUD/models"
	"localSandbox/SimpleCRUD/repository"
)

type CommentService interface {
	Save(models.Comment)
	SaveAll([]models.Comment)
	Update(models.Comment)
	FindAll() []models.Comment
	FindByID(int) models.Comment
	DeleteByID(int)
}

type CommentServiceImpl struct {
}

func (service CommentServiceImpl) Save(comment models.Comment) {
	repository.SaveToDB(&comment)
}

func (service CommentServiceImpl) SaveAll(comments []models.Comment) {
	repository.SaveToDB(&comments)
}

func (service CommentServiceImpl) Update(comment models.Comment) {
	var existingComment models.Comment
	repository.FindByID(&existingComment, comment.ID)
	if comment.Name != "" {
		existingComment.Name = comment.Name
	}
	if comment.Email != "" {
		existingComment.Email = comment.Email
	}
	if comment.Body != "" {
		existingComment.Body = comment.Body
	}
	repository.SaveToDB(&comment)
}

func (service CommentServiceImpl) DeleteByID(id int) {
	repository.DeleteFromDB(&models.Comment{}, id)
}

func (service CommentServiceImpl) FindAll() []models.Comment {
	var comments []models.Comment
	repository.FindAll(&comments, "Posts.Comments")
	return comments
}

func (service CommentServiceImpl) FindByID(id int) models.Comment {
	var comment models.Comment
	repository.FindByID(&comment, id, "Posts.Comments")
	return comment
}
