package dischargetype

import (
    "smrp/controller/setup/discharge_type"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/discharge-types", dischargetype.LookupList)
    api.Get("/discharge-types", dischargetype.List)
    api.Post("/discharge-types", dischargetype.SearchList)
    api.Post("/discharge-type", dischargetype.Create)
    api.Get("/discharge-type/:id", dischargetype.Edit)
    api.Put("/discharge-type/:id", dischargetype.Update)
    api.Delete("/discharge-type/:id", dischargetype.Delete)
}
