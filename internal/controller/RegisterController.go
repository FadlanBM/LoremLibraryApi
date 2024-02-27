package controller

import (
	"errors"
	"github.com/api/config"
	helper "github.com/api/internal/helpers"
	"github.com/api/internal/models"
	"github.com/api/internal/request"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"path/filepath"
	"strconv"
)

const (
	uploadPath = "C:\\Users\\fadlan\\Documents\\LoremLibrary\\Web\\public\\storage\\profile"
)

// RegisterController godoc
// @Tags Register
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Request"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/auth/register [post]
func RegisterController(c *fiber.Ctx) error {
	reqregister := new(request.RegisterRequest)

	// Parse request body
	if err := c.BodyParser(reqregister); err != nil {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	// Check if the borrower with the Google ID already exists
	var existingGoogle models.Borrowers
	if err := config.DB.First(&existingGoogle, "google_id = ?", reqregister.GoogleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			borrowers := models.Borrowers{
				Email:        reqregister.Email,
				Name:         reqregister.Name,
				Google_ID:    reqregister.GoogleID,
				Phone_Number: reqregister.PhoneNumber,
				Address:      reqregister.Address,
			}

			if err := config.DB.Create(&borrowers).Error; err != nil {
				return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
			}

			return c.Status(200).JSON(fiber.Map{"Status": "Insert", "Message": "Akun terdaftar", "Data": borrowers})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	existingGoogle.Email = reqregister.Email
	existingGoogle.Name = reqregister.Name
	existingGoogle.Phone_Number = reqregister.PhoneNumber
	existingGoogle.Address = reqregister.Address

	if err := config.DB.Save(&existingGoogle).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"Status": "Update", "Message": "Akun diperbarui", "Data": existingGoogle})
}

// RegisterAvatar godoc
// @Tags Profile Petugas
// @Accept json
// @Produce json
// @Param id path int true "petugas ID"
// @Param image formData file true "Image Upload"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/profile/petugas/image/{id} [put]
func RegisterAvatar(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var Borrowers models.Borrowers
	if err := config.DB.First(&Borrowers, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "false", "message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	targetData := Borrowers.Avatar
	filePath := filepath.Join(uploadPath, targetData)

	exists, err := helper.IsDataInDirectory(uploadPath, targetData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	if exists {
		if err := helper.RemoveFile(filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "message": err.Error()})
		}
	}

	filename, err := helper.HandleFileUpload(c, uploadPath, "image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	Borrowers.Avatar = filename

	if err := config.DB.Save(&Borrowers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "true", "message": "Peminjam image update"})
}
