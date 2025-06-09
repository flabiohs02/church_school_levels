package handler

import (
	"strconv"

	"church_school_levels/domain"
	"church_school_levels/usecase"

	"github.com/gofiber/fiber/v2"
)

type SchoolLevelHandler struct {
	slService *usecase.SchoolLevelService
}

func NewSchoolLevelHandler(service *usecase.SchoolLevelService) *SchoolLevelHandler {
	return &SchoolLevelHandler{slService: service}
}

func (h *SchoolLevelHandler) CreateSchoolLevel(c *fiber.Ctx) error {
	var sl domain.SchoolLevel
	if err := c.BodyParser(&sl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.slService.CreateSchoolLevel(&sl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(sl)
}

func (h *SchoolLevelHandler) GetSchoolLevelByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	sl, err := h.slService.GetSchoolLevelByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "School level not found"})
	}

	return c.Status(fiber.StatusOK).JSON(sl)
}

func (h *SchoolLevelHandler) GetAllSchoolLevels(c *fiber.Ctx) error {
	sls, err := h.slService.GetAllSchoolLevels()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(sls)
}

func (h *SchoolLevelHandler) UpdateSchoolLevel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var sl domain.SchoolLevel
	if err := c.BodyParser(&sl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sl.ID = uint(id)
	if err := h.slService.UpdateSchoolLevel(&sl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(sl)
}

func (h *SchoolLevelHandler) DeleteSchoolLevel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.slService.DeleteSchoolLevel(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
