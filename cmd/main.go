package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	connDB "github.com/BreakDown-CS/api-ecommerce/databases"
	"github.com/BreakDown-CS/api-ecommerce/internal/bootstrap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: false,
	}))

	_, err = connDB.PostgresConn()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	bootstrap.Bulid(app, connDB.ConnPostgres)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "4010"
	}
	serverAddress := fmt.Sprintf("0.0.0.0:%s", port)

	go func() {
		if err := app.Listen(serverAddress); err != nil {
			log.Printf("Server stopped: %v", err)
		}
	}()

	<-stop

	fmt.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}

	connDB.ClosePostgres()
	fmt.Println("Database connection closed")
}
