package relationship

import (
    "smrp/controller/setup/relationship"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/relationships", relationship.LookupList)
    api.Get("/relationships", relationship.List)
    api.Post("/relationships", relationship.SearchList)
    api.Post("/relationship", relationship.Create)
    api.Get("/relationship/:id", relationship.Edit)
    api.Put("/relationship/:id", relationship.Update)
    api.Delete("/relationship/:id", relationship.Delete)
}
