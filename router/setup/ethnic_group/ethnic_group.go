package ethnicgroup

import (
    "smrp/controller/setup/ethnic_group"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/ethnic-groups", ethnicgroup.LookupList)
    api.Get("/ethnic-groups", ethnicgroup.List)
    api.Post("/ethnic-groups", ethnicgroup.SearchList)
    api.Post("/ethnic-group", ethnicgroup.Create)
    api.Get("/ethnic-group/:id", ethnicgroup.Edit)
    api.Put("/ethnic-group/:id", ethnicgroup.Update)
    api.Delete("/ethnic-group/:id", ethnicgroup.Delete)
}
