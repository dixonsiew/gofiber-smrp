package personcategorycode

import (
    "smrp/controller/setup/person_category_code"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/person-category-codes", personcategorycode.LookupList)
    api.Get("/person-category-codes", personcategorycode.List)
    api.Post("/person-category-codes", personcategorycode.SearchList)
    api.Post("/person-category-code", personcategorycode.Create)
    api.Get("/person-category-code/:id", personcategorycode.Edit)
    api.Put("/person-category-code/:id", personcategorycode.Update)
    api.Delete("/person-category-code/:id", personcategorycode.Delete)
}
