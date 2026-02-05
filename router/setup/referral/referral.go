package referral

import (
    "smrp/controller/setup/referral"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api")
    api.Use(middleware.JWTProtected)
    api.Get("/lookup/referrals", referral.LookupList)
    api.Get("/referrals", referral.List)
    api.Post("/referrals", referral.SearchList)
    api.Post("/referral", referral.Create)
    api.Get("/referral/:id", referral.Edit)
    api.Put("/referral/:id", referral.Update)
    api.Delete("/referral/:id", referral.Delete)
}
