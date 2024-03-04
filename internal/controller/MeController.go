package controller

import (
	"errors"
	"github.com/api/config"
	"github.com/api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// MePeminjamController godoc
// @Tags Authorization
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseAuthSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/me [get]
func MePeminjamController(c *fiber.Ctx) error {
	userClaims, ok := c.Locals("borrowers").(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Failed get data"})
	}

	userID := userClaims["id"].(float64)

	var borrowers models.Borrowers
	if err := config.DB.First(&borrowers, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Success get data ", "data": borrowers})
}

// MeDeleteController godoc
// @Tags Authorization
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseAuthSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/auth/me/delete [get]
func MeDeleteController(c *fiber.Ctx) error {
	var reviews []models.Reviews
	var collection []models.Collection
	var listLending []models.ListLending
	var lending []models.Lending

	userClaims, ok := c.Locals("borrowers").(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Failed to get user data"})
	}

	userID := userClaims["id"].(float64)

	var borrowers models.Borrowers
	if err := config.DB.First(&borrowers, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data not found"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to fetch user data: " + err.Error()})
	}

	// Hapus koleksi yang terkait dengan pengguna
	if err := config.DB.Where("borrowers_id = ?", borrowers.ID).Unscoped().Delete(&collection).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete collection data: " + err.Error()})
	}

	// Hapus ulasan yang terkait dengan pengguna
	if err := config.DB.Where("borrowers_id = ?", borrowers.ID).Unscoped().Delete(&reviews).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete reviews data: " + err.Error()})
	}

	// Hapus peminjaman yang terkait dengan pengguna
	if err := config.DB.Where("borrowers_id = ?", borrowers.ID).Find(&lending).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to fetch lending data: " + err.Error()})
	}

	for _, v := range lending {
		// Hapus daftar peminjaman yang terkait dengan setiap peminjaman
		if err := config.DB.Where("lending_id = ?", v.ID).Unscoped().Delete(&listLending).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete list lending data: " + err.Error()})
		}
	}

	// Hapus peminjaman yang terkait dengan pengguna
	if err := config.DB.Where("borrowers_id = ?", borrowers.ID).Unscoped().Delete(&lending).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete lending data: " + err.Error()})
	}

	// Hapus pengguna
	if err := config.DB.Unscoped().Delete(&borrowers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete user data: " + err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Data deleted successfully"})

}
