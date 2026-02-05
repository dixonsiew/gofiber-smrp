package title

import (
    "smrp/controller/setup/title"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/titles", title.LookupList)
    api.Get("/titles", title.List)
    api.Post("/titles", title.SearchList)
    api.Post("/title", title.Create)
    api.Get("/title/:id", title.Edit)
    api.Put("/title/:id", title.Update)
    api.Delete("/title/:id", title.Delete)
}
