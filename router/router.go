package router

// https://www.google.com/search?q=httperrorinterceptor+depends+on+app_initializer&oq=httperrorinterceptor+depends+on+app_initializer&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIHCAEQABjvBTIHCAIQABjvBdIBCTIwOTg1ajBqN6gCCLACAfEF1DF3FC55WYs&sourceid=chrome&ie=UTF-8

import (
	"smrp/router/auth"
    "smrp/router/report"
	"smrp/router/setup"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/smrp")
    auth.SetupRoutes(api)
    report.SetupRoutes(api)
    setup.SetupRoutes(api)
}
