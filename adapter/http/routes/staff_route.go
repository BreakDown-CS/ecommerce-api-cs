package router

import (
	"github.com/BreakDown-CS/api-ecommerce/adapter/http/handler"
	"github.com/gofiber/fiber/v2"
)

func RegisterStaffRoutes(app *fiber.App, staffHandler *handler.StaffHandler) {
	staff := app.Group("/api/staff")

	staff.Post("/getStaffList", staffHandler.GetStaffList)
}
