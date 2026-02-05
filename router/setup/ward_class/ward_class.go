package wardclass

import (
    "smrp/controller/setup/ward_class"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/ward-classes", wardclass.LookupList)
    api.Get("/ward-classes", wardclass.List)
    api.Post("/ward-classes", wardclass.SearchList)
    api.Post("/ward-class", wardclass.Create)
    api.Get("/ward-class/:id", wardclass.Edit)
    api.Put("/ward-class/:id", wardclass.Update)
    api.Delete("/ward-class/:id", wardclass.Delete)
}
