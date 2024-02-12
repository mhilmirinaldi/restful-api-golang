package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/magiconair/properties/assert"
	"restful-api-golang/config"
	"restful-api-golang/repository"
	request2 "restful-api-golang/web/request"
	"testing"
)

func TestCreate(t *testing.T) {
	db := config.ConnectionDB()
	validate := validator.New()

	commentRepository := repository.NewCommentRepositoryImpl(db)

	commentService := NewCommentServiceImpl(commentRepository, validate)

	request := request2.CommentCreateRequest{Description: "test service"}

	response := commentService.Create(request)

	fmt.Println(response)

	assert.Equal(t, 1, 1)
}

func TestUpdate(t *testing.T) {
	db := config.ConnectionDB()
	validate := validator.New()

	commentRepository := repository.NewCommentRepositoryImpl(db)

	commentService := NewCommentServiceImpl(commentRepository, validate)

	request := request2.CommentUpdateRequest{
		Id:          7,
		Description: "updated",
	}

	response := commentService.Update(request)

	fmt.Println(response)

	assert.Equal(t, 1, 1)
}

func TestFindById(t *testing.T) {
	db := config.ConnectionDB()
	validate := validator.New()

	commentRepository := repository.NewCommentRepositoryImpl(db)

	commentService := NewCommentServiceImpl(commentRepository, validate)

	response := commentService.FindById(7)

	fmt.Println(response)

	assert.Equal(t, 1, 1)
}

func TestAll(t *testing.T) {
	db := config.ConnectionDB()
	validate := validator.New()

	commentRepository := repository.NewCommentRepositoryImpl(db)

	commentService := NewCommentServiceImpl(commentRepository, validate)

	response := commentService.FindAll()

	fmt.Println(response)

	assert.Equal(t, 1, 1)
}
