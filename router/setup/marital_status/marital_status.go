package maritalstatus

import (
    "smrp/controller/setup/marital_status"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/marital-statuses", maritalstatus.LookupList)
    api.Get("/marital-statuses", maritalstatus.List)
    api.Post("/marital-statuses", maritalstatus.SearchList)
    api.Post("/marital-status", maritalstatus.Create)
    api.Get("/marital-status/:id", maritalstatus.Edit)
    api.Put("/marital-status/:id", maritalstatus.Update)
    api.Delete("/marital-status/:id", maritalstatus.Delete)
}
