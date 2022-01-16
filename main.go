package main

import (
	"golang-rest-api/pkg/configs"
	"golang-rest-api/pkg/middleware"
	"golang-rest-api/pkg/routes"
	"golang-rest-api/utils"

	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env
)

func main() {

	// fiber config
	config := configs.FiberConfig()

	// define fiber app with config
	app := fiber.New(config)

	// enable middleware
	middleware.FiberMiddleware(app)

	// add routes
	routes.PublicRoutes(app)

	// init server
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
