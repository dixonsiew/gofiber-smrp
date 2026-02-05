package income

import (
    "smrp/controller/setup/income"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/incomes", income.LookupList)
    api.Get("/incomes", income.List)
    api.Post("/incomes", income.SearchList)
    api.Post("/income", income.Create)
    api.Get("/income/:id", income.Edit)
    api.Put("/income/:id", income.Update)
    api.Delete("/income/:id", income.Delete)
}
