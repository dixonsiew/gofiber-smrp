package religion

import (
    "smrp/controller/setup/religion"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/religions", religion.LookupList)
    api.Get("/religions", religion.List)
    api.Post("/religions", religion.SearchList)
    api.Post("/religion", religion.Create)
    api.Get("/religion/:id", religion.Edit)
    api.Put("/religion/:id", religion.Update)
    api.Delete("/religion/:id", religion.Delete)
}
