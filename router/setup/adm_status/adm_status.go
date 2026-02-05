package admstatus

import (
    "smrp/controller/setup/adm_status"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/adm-statuses", admstatus.LookupList)
    api.Get("/adm-statuses", admstatus.List)
    api.Post("/adm-statuses", admstatus.SearchList)
    api.Post("/adm-status", admstatus.Create)
    api.Get("/adm-status/:id", admstatus.Edit)
    api.Put("/adm-status/:id", admstatus.Update)
    api.Delete("/adm-status/:id", admstatus.Delete)
}