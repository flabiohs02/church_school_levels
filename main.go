package main

import (
	"log"

	"church_school_levels/config"
	"church_school_levels/handler"
	"church_school_levels/repository"
	"church_school_levels/routers"
	"church_school_levels/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := config.InitDb()

	// Dependency Injection
	slRepo := repository.NewGormSchoolLevelRepository(db)
	slService := usecase.NewSchoolLevelService(slRepo)
	slHandler := handler.NewSchoolLevelHandler(slService)

	// Set up Fiber app
	app := fiber.New()

	// Setup routes
	routers.SetupSchoolLevelRoutes(app, slHandler)

	log.Println("Server listening on :8181")
	if err := app.Listen(":8181"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
