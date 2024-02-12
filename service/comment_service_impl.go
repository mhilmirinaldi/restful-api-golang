package service

import (
	"github.com/go-playground/validator/v10"
	"restful-api-golang/helper"
	"restful-api-golang/model"
	"restful-api-golang/repository"
	"restful-api-golang/web/request"
	"restful-api-golang/web/response"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	Validate          *validator.Validate
}

func NewCommentServiceImpl(commentRepository repository.CommentRepository, Validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		Validate:          Validate,
	}
}

func (c *CommentServiceImpl) Create(request request.CommentCreateRequest) response.CommentResponse {
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	comment := model.Comment{
		Description: request.Description,
	}

	result := c.CommentRepository.Save(comment)

	commentResponse := response.CommentResponse{
		Id:          result.ID,
		Description: result.Description,
	}

	return commentResponse
}

func (c *CommentServiceImpl) Update(request request.CommentUpdateRequest) response.CommentResponse {
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	comment := model.Comment{
		ID:          request.Id,
		Description: request.Description,
	}

	result := c.CommentRepository.Update(comment)

	commentResponse := response.CommentResponse{
		Id:          result.ID,
		Description: result.Description,
	}

	return commentResponse

}

func (c *CommentServiceImpl) Delete(commentId int) {
	c.CommentRepository.Delete(commentId)
}

func (c *CommentServiceImpl) FindById(commentId int) response.CommentResponse {
	comment, err := c.CommentRepository.FindById(commentId)
	helper.PanicIfError(err)

	return response.CommentResponse{
		Id:          comment.ID,
		Description: comment.Description,
	}
}

func (c *CommentServiceImpl) FindAll() []response.CommentResponse {
	result := c.CommentRepository.FindAll()
	var comments []response.CommentResponse

	for _, value := range result {
		comment := response.CommentResponse{
			Id:          value.ID,
			Description: value.Description,
		}
		comments = append(comments, comment)
	}

	return comments
}
