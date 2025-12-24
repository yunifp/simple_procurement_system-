package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"procurement-system/internal/handlers"
	"procurement-system/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")


	app.Get("/swagger/*", swagger.HandlerDefault)


	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	v1 := api.Group("/v1", middleware.Protected())


	v1.Get("/suppliers", handlers.GetSuppliers)
	v1.Get("/suppliers/:id", handlers.GetSupplier)
	v1.Post("/suppliers", handlers.CreateSupplier)
	v1.Put("/suppliers/:id", handlers.UpdateSupplier)
	v1.Delete("/suppliers/:id", handlers.DeleteSupplier)


	v1.Get("/items", handlers.GetItems)
	v1.Get("/items/:id", handlers.GetItem)
	v1.Post("/items", handlers.CreateItem)
	v1.Put("/items/:id", handlers.UpdateItem)
	v1.Delete("/items/:id", handlers.DeleteItem)


	v1.Post("/purchasing", handlers.CreatePurchase)
}