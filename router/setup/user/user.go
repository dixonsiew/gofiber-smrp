package user

import (
    "smrp/controller/setup/user"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/users", user.List)
    api.Post("/users", user.SearchList)
    api.Post("/user", user.Create)
    api.Get("/user/:id", user.Edit)
    api.Put("/user/:id", user.Update)
    api.Delete("/user/:id", user.Delete)
}
