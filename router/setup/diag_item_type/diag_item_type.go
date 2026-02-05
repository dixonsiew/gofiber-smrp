package diagitemtype

import (
    "smrp/controller/setup/diag_item_type"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/diag-item-types", diagitemtype.LookupList)
    api.Get("/diag-item-types", diagitemtype.List)
    api.Post("/diag-item-types", diagitemtype.SearchList)
    api.Post("/diag-item-type", diagitemtype.Create)
    api.Get("/diag-item-type/:id", diagitemtype.Edit)
    api.Put("/diag-item-type/:id", diagitemtype.Update)
    api.Delete("/diag-item-type/:id", diagitemtype.Delete)
}
