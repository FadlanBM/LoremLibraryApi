package controller

import (
	"errors"
	"github.com/api/config"
	"github.com/api/internal/models"
	"github.com/api/internal/request"
	response "github.com/api/internal/respose"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
)

// CommentCreate godoc
// @Tags Comment
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Param request body request.CommentRequest true "Request"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/comment [post]
func CommentCreate(c *fiber.Ctx) error {
	req := new(request.CommentRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

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

	var books models.Books
	if err := config.DB.First(&books, req.BooksID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	reviews := models.Reviews{
		BorrowersID: borrowers.ID,
		BooksID:     books.ID,
		Rating:      req.Rating,
		Messages:    req.Messages,
	}

	if err := config.DB.Create(&reviews).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Berhasil melakukan transaksi"})
}

// CommentIndex godoc
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/comment/{id} [get]
func CommentIndex(c *fiber.Ctx) error {
	idParam := c.Params("id")
	var book models.Books
	var reviews []models.Reviews
	var borrowers models.Borrowers
	var res []response.RatingsResponse

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "ID harus berupa angka"})
	}

	if err := config.DB.Find(&book, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	if err := config.DB.Find(&reviews, "books_id = ?", book.ID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	for _, v := range reviews {

		if err := config.DB.First(&borrowers, v.BorrowersID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
			}
			return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
		}
		ratingsResponse := response.RatingsResponse{
			ID:           v.ID,
			BorrowerName: borrowers.Name,
			Message:      v.Messages,
			Ratings:      v.Rating,
		}
		res = append(res, ratingsResponse)
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Succcess get data", "data": res})
}

// CommentCheck godoc
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/comment/check/{id} [get]
func CommentCheck(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "ID harus berupa angka"})
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	userClaims, ok := c.Locals("borrowers").(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Failed get data"})
	}
	userID := int(userClaims["id"].(float64)) // Konversi float64 ke int

	var borrower models.Borrowers
	if err := config.DB.First(&borrower, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var review models.Reviews
	if err := config.DB.Where("books_id = ? AND borrowers_id = ?", book.ID, borrower.ID).First(&review).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "true", "message": "Borrowers Have Not Commented"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Borrowers Have Commented"})
}
