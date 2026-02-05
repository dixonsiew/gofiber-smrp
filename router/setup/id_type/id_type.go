package idtype

import (
    "smrp/controller/setup/id_type"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/id-types", idtype.LookupList)
    api.Get("/id-types", idtype.List)
    api.Post("/id-types", idtype.SearchList)
    api.Post("/id-type", idtype.Create)
    api.Get("/id-type/:id", idtype.Edit)
    api.Put("/id-type/:id", idtype.Update)
    api.Delete("/id-type/:id", idtype.Delete)
}
