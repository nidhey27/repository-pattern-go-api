package server

import (
	"flag"
	"net/http"
	"os"
	"rest-api-redis/internal/handlers"
	"rest-api-redis/pkg/database"
	"rest-api-redis/pkg/repository"
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
	repository := repository.InitRepository(db)
	h := handlers.InitHandler(repository)

	app.Get("/", func(c *fiber.Ctx) error {
		return utils.SendResponse(http.StatusOK, "Users Management using Repository Pattern by Nidhey Indurkar", "", make([]string, 0), c)
	})

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return utils.SendResponse(http.StatusOK, "Healthy", "", make([]string, 0), c)
	})

	app.Post("/api/user", h.UserHandler.CreateUser)
	app.Get("/api/user/:id", h.UserHandler.GetUser)
	app.Get("/api/user", h.UserHandler.GetUsers)
	app.Delete("/api/user/:id", h.UserHandler.DeleteUser)
	app.Put("/api/user/:id", h.UserHandler.UpdateUser)

	return app.Listen(*addr)
}
