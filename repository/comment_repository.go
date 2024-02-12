package repository

import "restful-api-golang/model"

type CommentRepository interface {
	Save(comment model.Comment) model.Comment
	Update(comment model.Comment) model.Comment
	Delete(commentId int)
	FindById(commentId int) (model.Comment, error)
	FindAll() []model.Comment
}
