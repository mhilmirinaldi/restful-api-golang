package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"restful-api-golang/config"
	"restful-api-golang/controller"
	"restful-api-golang/repository"
	"restful-api-golang/router"
	"restful-api-golang/service"
)

func main() {
	db := config.ConnectionDB()
	validate := validator.New()

	commentRepository := repository.NewCommentRepositoryImpl(db)

	commentService := service.NewCommentServiceImpl(commentRepository, validate)

	commentController := controller.NewCommentControllerImpl(commentService)

	routes := router.NewRouter(commentController)

	app := fiber.New()

	app.Mount("", routes)
	log.Fatal(app.Listen(":8000"))
}
