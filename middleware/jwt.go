package middleware

import (
    "smrp/model"
    userService "smrp/service/user"
    tokenService "smrp/service/token"
    "smrp/utils"

    jwtware "github.com/gofiber/contrib/jwt"
    "github.com/gofiber/fiber/v2"
)

func JWTProtected(c *fiber.Ctx) error {
    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(utils.JWT_SECRET)},
        ContextKey: "jwt",
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "statusCode": fiber.StatusUnauthorized,
                "message":    "Unauthorized",
            })
        },
    })(c)
}

func ValidateToken(c *fiber.Ctx) (int, *model.User, error) {
    _, id, err := tokenService.DecodeToken(c)
    if err != nil {
        return id, nil, err
    }

    user, err := userService.FindById(id)
    if err != nil || user == nil {
        return id, user, err
    }

    return id, user, nil
}

func NoContent(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
        "statusCode": fiber.StatusNoContent,
        "message":    "",
    })
}

func Unauthorized(c *fiber.Ctx) error {
    mx := fiber.Map{
        "statusCode": fiber.StatusUnauthorized,
        "message":    "User Not Found",
    }
    return c.Status(fiber.StatusUnauthorized).JSON(mx)
}