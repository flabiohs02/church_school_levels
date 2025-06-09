package routers

import (
	"church_school_levels/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupSchoolLevelRoutes(app *fiber.App, slHandler *handler.SchoolLevelHandler) {
	v1 := app.Group("/api/v1")
	{
		v1.Post("/school-levels", slHandler.CreateSchoolLevel)
		v1.Get("/school-levels/:id", slHandler.GetSchoolLevelByID)
		v1.Get("/school-levels", slHandler.GetAllSchoolLevels)
		v1.Put("/school-levels/:id", slHandler.UpdateSchoolLevel)
		v1.Delete("/school-levels/:id", slHandler.DeleteSchoolLevel)
	}
}
