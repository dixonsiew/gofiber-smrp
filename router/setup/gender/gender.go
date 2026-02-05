package gender

import (
    "smrp/controller/setup/gender"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/genders", gender.LookupList)
    api.Get("/genders", gender.List)
    api.Post("/genders", gender.SearchList)
    api.Post("/gender", gender.Create)
    api.Get("/gender/:id", gender.Edit)
    api.Put("/gender/:id", gender.Update)
    api.Delete("/gender/:id", gender.Delete)
}
