package router

import (
	"github.com/gofiber/fiber/v2"
	"restful-api-golang/controller"
)

func NewRouter(controller *controller.CommentControllerImpl) *fiber.App {
	router := fiber.New()

	router.Route("/comments", func(router fiber.Router) {
		router.Get("", controller.FindAll)
		router.Post("", controller.Create)
	})

	router.Route("/comments/:commentId", func(router fiber.Router) {
		router.Get("", controller.FindById)
		router.Put("", controller.Update)
		router.Delete("", controller.Delete)
	})

	return router
}
