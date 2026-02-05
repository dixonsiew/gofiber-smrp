package state

import (
    "smrp/controller/setup/state"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/states", state.LookupList)
    api.Get("/states", state.List)
    api.Post("/states", state.SearchList)
    api.Post("/state", state.Create)
    api.Get("/state/:id", state.Edit)
    api.Put("/state/:id", state.Update)
    api.Delete("/state/:id", state.Delete)
}
