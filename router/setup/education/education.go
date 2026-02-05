package education

import (
    "smrp/controller/setup/education"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/educations", education.LookupList)
    api.Get("/educations", education.List)
    api.Post("/educations", education.SearchList)
    api.Post("/education", education.Create)
    api.Get("/educations/:id", education.Edit)
    api.Put("/educations/:id", education.Update)
    api.Delete("/educations/:id", education.Delete)
}
