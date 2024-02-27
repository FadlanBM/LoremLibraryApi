package controller

import (
	"github.com/api/config"
	helper "github.com/api/internal/helpers"
	"github.com/api/internal/models"
	"github.com/api/internal/request"
	"github.com/gofiber/fiber/v2"
)

// AuthController godoc
// @Tags Authorization
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Param request body request.AuthRequest true "Request"
// @Success 200 {object} response.ResponseAuthSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/auth [post]
func AuthController(c *fiber.Ctx) error {
	req := new(request.AuthRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "Error", "message": err.Error()})
	}

	borrowers := new(models.Borrowers)
	if err := config.DB.First(borrowers, "google_id = ?", req.GoogleID).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "Error", "message": "Account Not Found "})
	}

	token, err := helper.GenerateToken(borrowers)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "Token Generation Failed"})
	}

	return c.Status(200).JSON(fiber.Map{"token": token, "role": borrowers.Role, "active": borrowers.Active})
}
