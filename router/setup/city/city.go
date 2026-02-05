package city

import (
    "smrp/controller/setup/city"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/cities", city.LookupList)
    api.Get("/cities", city.List)
    api.Post("/cities", city.SearchList)
    api.Post("/city", city.Create)
    api.Get("/city/:id", city.Edit)
    api.Put("/city/:id", city.Update)
    api.Delete("/city/:id", city.Delete)
}
