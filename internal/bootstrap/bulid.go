package bootstrap

import (
	"github.com/BreakDown-CS/api-ecommerce/adapter/http/handler"
	router "github.com/BreakDown-CS/api-ecommerce/adapter/http/routes"
	"github.com/BreakDown-CS/api-ecommerce/adapter/persistance/repositories"
	"github.com/BreakDown-CS/api-ecommerce/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Bulid(app *fiber.App, DBPostgres *gorm.DB) {
	// Staff
	staffRepositories := repositories.NewStaffRepositories(DBPostgres)
	staffServices := service.NewStaffService(staffRepositories)
	staffHandler := handler.NewStaffHandler(staffServices)
	router.RegisterStaffRoutes(app, staffHandler)
}
