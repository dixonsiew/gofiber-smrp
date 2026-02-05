package occupation

import (
    "smrp/controller/setup/occupation"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/occupations", occupation.LookupList)
    api.Get("/occupations", occupation.List)
    api.Post("/occupations", occupation.SearchList)
    api.Post("/occupation", occupation.Create)
    api.Get("/occupation/:id", occupation.Edit)
    api.Put("/occupation/:id", occupation.Update)
    api.Delete("/occupation/:id", occupation.Delete)
}
