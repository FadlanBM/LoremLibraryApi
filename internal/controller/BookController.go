package controller

import (
	"github.com/api/config"
	"github.com/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// IndexBook godoc
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/book [get]
func BookIndex(c *fiber.Ctx) error {
	var book []models.Books

	if err := config.DB.Find(&book).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Succcess get data", "data": book})
}
