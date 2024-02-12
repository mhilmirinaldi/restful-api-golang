package service

import (
	"restful-api-golang/web/request"
	"restful-api-golang/web/response"
)

type CommentService interface {
	Create(request request.CommentCreateRequest) response.CommentResponse
	Update(updateRequest request.CommentUpdateRequest) response.CommentResponse
	Delete(commentId int)
	FindById(commentId int) response.CommentResponse
	FindAll() []response.CommentResponse
}
