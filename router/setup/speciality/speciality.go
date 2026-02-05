package speciality

import (
    "smrp/controller/setup/speciality"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/specialities", speciality.LookupList)
    api.Get("/specialities", speciality.List)
    api.Post("/specialities", speciality.SearchList)
    api.Post("/speciality", speciality.Create)
    api.Get("/speciality/:id", speciality.Edit)
    api.Put("/speciality/:id", speciality.Update)
    api.Delete("/speciality/:id", speciality.Delete)
}
