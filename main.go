package main

// https://betterstack.com/community/guides/scaling-go/postgresql-pgx-golang/

import (
    "errors"
    "fmt"
    "html/template"
    "os"
    "smrp/config"
    "smrp/database"
    _ "smrp/docs"
    "smrp/router"
    "smrp/utils"
    "strings"

    "github.com/gofiber/contrib/fiberzerolog"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/compress"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/monitor"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/swagger"
    redoc "github.com/natebwangsut/fiber-redoc"
)

// @title Swagger SMRP Backend API
// @version 1.0
// @description SMRP Backend API description.
// @BasePath /smrp
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @type http
// @scheme bearer
// @bearerFormat JWT
// @description Type "Bearer" followed by a space and JWT token.
func main() {
    defer utils.CatchPanic("main")
    runLogFile, _ := os.OpenFile(
        "app.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
    defer runLogFile.Close()
    utils.SetValidator()
    utils.SetLogger(runLogFile)
    port := config.Config("port")
    app := fiber.New(fiber.Config{
        Prefork: false,
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            var e *fiber.Error
            if errors.As(err, &e) {
                code = e.Code
            }

            return c.Status(code).JSON(fiber.Map{
                "statusCode": code,
                "message":    err.Error(),
            })
        },
    })
    app.Use(recover.New())
    app.Use(compress.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins:  "*",
        ExposeHeaders: strings.Join([]string{"Authorization", "filename", utils.X_TOTAL_COUNT, utils.X_TOTAL_PAGE}, ","),
    }))
    app.Use(fiberzerolog.New(fiberzerolog.Config{
        Logger: &utils.Logger,
    }))
    // app.Use(jwtware.New(jwtware.Config{
    //     SigningKey: jwtware.SigningKey{Key: []byte(utils.JWT_SECRET)},
    // }))
    database.ConnectDB()
    database.ConnectDBRs()
    database.ConnectMongo()
    defer database.CloseDB()
    defer database.CloseDBRs()
    defer database.CloseMongo()

    basePath := "smrp"
    initSwagger(app, basePath)
    app.Get("/smrp/metrics", monitor.New())
    router.SetupRoutes(app)

    err := app.Listen(fmt.Sprintf(":%s", port))

    if err != nil {
        utils.Logger.Fatal().Err(err).Msg("Fiber app error")
    }
}

func initSwagger(app *fiber.App, basePath string) {
    b, _ := os.ReadFile("./public/css/theme-flattop.css")
    css := string(b)

    cfg := swagger.Config{
        URL:          "doc.json",
        DeepLinking:  true,
        DocExpansion: "list",
        Title:        "Swagger SMRP Backend API",
        SyntaxHighlight: &swagger.SyntaxHighlightConfig{
            Activate: true,
            Theme:    "arta",
        },
        CustomStyle:          template.CSS(css),
        PersistAuthorization: true,
    }

    app.Get(fmt.Sprintf("/%s/docs/*", basePath), swagger.New(cfg))
    app.Get(fmt.Sprintf("/%s/redocs/*", basePath), redoc.Handler)

    app.Static(fmt.Sprintf("/%s/static", basePath), "./public")
}
