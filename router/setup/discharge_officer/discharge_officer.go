package dischargeofficer

import (
    "smrp/controller/setup/discharge_officer"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/discharge-officers", dischargeofficer.LookupList)
    api.Get("/discharge-officers", dischargeofficer.List)
    api.Post("/discharge-officers", dischargeofficer.SearchList)
    api.Post("/discharge-officer", dischargeofficer.Create)
    api.Get("/discharge-officer/:id", dischargeofficer.Edit)
    api.Put("/discharge-officer/:id", dischargeofficer.Update)
    api.Delete("/discharge-officer/:id", dischargeofficer.Delete)
}
