package deliverytype

import (
    "smrp/controller/setup/delivery_type"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/delivery-types", deliverytype.LookupList)
    api.Get("/delivery-types", deliverytype.List)
    api.Post("/delivery-types", deliverytype.SearchList)
    api.Post("/delivery-type", deliverytype.Create)
    api.Get("/delivery-type/:id", deliverytype.Edit)
    api.Put("/delivery-type/:id", deliverytype.Update)
    api.Delete("/delivery-type/:id", deliverytype.Delete)
}
