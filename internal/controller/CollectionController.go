package controller

import (
	"errors"
	"github.com/api/config"
	"github.com/api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
)

// CollectionIndex godoc
// @Tags CollectionController
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/collection/{id} [get]
func CollectionIndex(c *fiber.Ctx) error {
	var collection []models.Collection
	userClaims, ok := c.Locals("borrowers").(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Failed get data"})
	}

	userID := userClaims["id"].(float64)

	if err := config.DB.Find(&collection, userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Success get data", "data": collection})
}

// CollectionCreate godoc
// @Tags CollectionController
// @Accept json
// @Produce json
// @Param request body request.CollectionRequest true "Request"
// @Success 200 {object} response.ResponseSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/collection/{id} [post]
func CollectionCreate(c *fiber.Ctx) error {
	var collection models.Collection

	userClaims, ok := c.Locals("borrowers").(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Gagal mendapatkan data"})
	}

	userID := uint(userClaims["id"].(float64))

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	collection = models.Collection{
		BorrowersID: userID,
		BooksID:     uint(id),
	}

	if err := config.DB.Create(&collection).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Data admin berhasil dimasukkan"})
}

// CollectionDelete godoc
// @Tags CollectionController
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/collection [delete]
func CollectionDelete(c *fiber.Ctx) error {
	userClaims, ok := c.Locals("borrowers").(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Gagal mendapatkan data"})
	}

	userID := uint(userClaims["id"].(float64))

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	var collection models.Collection
	if err := config.DB.First(&collection, "borrowers_id = ? AND books_id = ?", userID, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	if res := config.DB.Delete(&collection); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": res.Error.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Data berhasil dihapus"})

}
