package report

import (
    "smrp/router/report/master_pd101"
    "smrp/router/report/master_pd102"
    "smrp/router/report/master_pd105"
    "smrp/router/report/master_pd301"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    masterpd101.SetupRoutes(router)
    masterpd102.SetupRoutes(router)
    masterpd105.SetupRoutes(router)
    masterpd301.SetupRoutes(router)
}
