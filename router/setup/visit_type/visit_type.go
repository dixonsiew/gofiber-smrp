package visittype

import (
    "smrp/controller/setup/visit_type"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/visit-types", visittype.LookupList)
    api.Get("/visit-types", visittype.List)
    api.Post("/visit-types", visittype.SearchList)
    api.Post("/visit-type", visittype.Create)
    api.Get("/visit-type/:id", visittype.Edit)
    api.Put("/visit-type/:id", visittype.Update)
    api.Delete("/visit-type/:id", visittype.Delete)
}
