package router

import (
	"github.com/abdukarimxalilov/demo-fiber-postgres/handler"
	"github.com/abdukarimxalilov/demo-fiber-postgres/middleware"
	"github.com/gofiber/fiber"
)

func SetupRoutes (app *fiber.App) { 
    api := app.Group("/api", middleware.AuthReq())  

    api.Get("/", handler.GetAllProducts)
    api.Get("/:id", handler.GetSingleProduct)
    api.Post("/", handler.CreateProduct)
    api.Delete("/:id", handler.DeleteProduct)
}