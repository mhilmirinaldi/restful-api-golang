package controller

import (
	"github.com/gofiber/fiber/v2"
	"restful-api-golang/helper"
	"restful-api-golang/service"
	"restful-api-golang/web/request"
	"restful-api-golang/web/response"
	"strconv"
)

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentControllerImpl(commentService service.CommentService) *CommentControllerImpl {
	return &CommentControllerImpl{CommentService: commentService}
}

func (c CommentControllerImpl) Create(ctx *fiber.Ctx) error {
	createRequest := request.CommentCreateRequest{}
	err := ctx.BodyParser(&createRequest)
	helper.PanicIfError(err)

	result := c.CommentService.Create(createRequest)

	webResponse := response.WebResponse{
		Code:    201,
		Status:  "Created",
		Message: "Comment successfully created",
		Data:    result,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c CommentControllerImpl) Update(ctx *fiber.Ctx) error {
	updateRequest := request.CommentUpdateRequest{}
	err := ctx.BodyParser(&updateRequest)
	helper.PanicIfError(err)

	commentId := ctx.Params("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)

	updateRequest.Id = id

	result := c.CommentService.Update(updateRequest)

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Comment successfully updated",
		Data:    result,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c CommentControllerImpl) Delete(ctx *fiber.Ctx) error {
	commentId := ctx.Params("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)

	c.CommentService.Delete(id)

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Comment successfully deleted",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func (c CommentControllerImpl) FindById(ctx *fiber.Ctx) error {
	commentId := ctx.Params("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)

	comment := c.CommentService.FindById(id)

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Comment successfully fetched by id",
		Data:    comment,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c CommentControllerImpl) FindAll(ctx *fiber.Ctx) error {
	comments := c.CommentService.FindAll()

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Comment successfully fetched",
		Data:    comments,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
