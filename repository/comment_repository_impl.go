package repository

import (
	"errors"
	"gorm.io/gorm"
	"restful-api-golang/helper"
	"restful-api-golang/model"
	"restful-api-golang/web/request"
)

type CommentRepositoryImpl struct {
	Db *gorm.DB
}

func NewCommentRepositoryImpl(Db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{Db: Db}

}

func (c *CommentRepositoryImpl) Save(comment model.Comment) model.Comment {
	result := c.Db.Create(&comment)
	helper.PanicIfError(result.Error)
	return comment
}

func (c *CommentRepositoryImpl) Update(comment model.Comment) model.Comment {
	var commentUpdated = request.CommentUpdateRequest{
		Id:          comment.ID,
		Description: comment.Description,
	}

	result := c.Db.Model(&comment).Updates(commentUpdated)
	helper.PanicIfError(result.Error)

	return comment
}

func (c *CommentRepositoryImpl) Delete(commentId int) {
	result := c.Db.Where("id = ?", commentId).Delete(&model.Comment{})
	helper.PanicIfError(result.Error)
}

func (c *CommentRepositoryImpl) FindById(commentId int) (model.Comment, error) {
	comment := model.Comment{}
	result := c.Db.Take(&comment, "id = ?", commentId)
	if result != nil {
		return comment, nil
	} else {
		// HANDLE ERROR IF NOT FOUND
		return comment, errors.New("comment not found")
	}
}

func (c *CommentRepositoryImpl) FindAll() []model.Comment {
	var comments []model.Comment
	result := c.Db.Find(&comments)
	helper.PanicIfError(result.Error)
	return comments

}
