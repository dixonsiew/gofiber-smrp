package country

import (
    "smrp/controller/setup/country"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/countries", country.LookupList)
    api.Get("/countries", country.List)
    api.Post("/countries", country.SearchList)
    api.Post("/country", country.Create)
    api.Get("/country/:id", country.Edit)
    api.Put("/country/:id", country.Update)
    api.Delete("/country/:id", country.Delete)
}
