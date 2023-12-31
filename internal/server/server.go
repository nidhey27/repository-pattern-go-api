package server

import (
	"flag"
	"net/http"
	"os"
	"rest-api-redis/internal/handlers"
	"rest-api-redis/pkg/database"
	"rest-api-redis/pkg/models"
	"rest-api-redis/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var addr = flag.String("addr", ":"+os.Getenv("PORT"), "")

func Run() error {
	flag.Parse()
	app := fiber.New()
	if *addr == ":" {
		*addr = ":8080"
	}

	app.Use(logger.New())
	app.Use(cors.New())

	db := database.GetDB()
	defer db.Close()

	db.Begin().AutoMigrate(&models.User{})

	app.Get("/", func(c *fiber.Ctx) error {
		return utils.SendResponse(http.StatusOK, "Movies & Web Series Library by Nidhey Indurkar", "", make([]string, 0), c)
	})

	app.Get("/healthchecker", func(c *fiber.Ctx) error {
		return utils.SendResponse(http.StatusOK, "Welcome to Golang, Fiber, and GORM", "", make([]string, 0), c)
	})

	app.Post("/api/user", handlers.CreateUser)
	app.Get("/api/user/:id", handlers.GetUser)
	app.Get("/api/user", handlers.GetUsers)
	app.Delete("/api/user/:id", handlers.DeleteUser)
	app.Put("/api/user/:id", handlers.UpdateUser)

	return app.Listen(*addr)
}
