package masterpd101

import (
    rp "smrp/controller/report/master_pd101"
    "smrp/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/api/master-pd101")
    api.Use(middleware.JWTProtected)
    api.Get("/export/rpt2", rp.JsonRH101)
    api.Get("/export/rpt1", rp.JsonPD101)
    api.Get("/export/rpt1/xlsx", rp.Xlsx)
    api.Get("/rpt1", rp.List)
    api.Post("/rpt1", rp.SearchList)
    api.Get("/rpt1/:id", rp.Edit)
    api.Put("/rpt1/:id", rp.Update)
}
