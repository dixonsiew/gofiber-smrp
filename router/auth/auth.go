package auth

import (
    "smrp/controller/auth"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    router.Post("/o/token", auth.Login)
    router.Post("/o/refresh-token", middleware.JWTProtected, auth.Refresh)
    router.Get("/api/current-user", middleware.JWTProtected, auth.UserDetails)
    router.Post("/api/change-password", middleware.JWTProtected, auth.ChangePassword)
}
