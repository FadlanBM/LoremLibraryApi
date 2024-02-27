package controller

import (
	"errors"
	"github.com/api/config"
	"github.com/api/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
	"strconv"
)

const (
	sampulPath = "E:\\UKK\\LoremLibrary\\storage\\app\\public\\sampul\\"
)

// BookIndex godoc
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

// BookDetail godoc
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/book/{id} [get]
func BookDetail(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Success ambil data ", "data": book})
}

// CategoryDetail godoc
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Failure 404 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
// @Router /api/book/category/{id} [get]
func CategoryDetail(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var book models.Books
	if err := config.DB.Preload("Book_Category").First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	// Menyiapkan array untuk menyimpan kategori buku
	var categories []models.Categories

	// Memuat kategori buku menggunakan Preload dengan ManyToMany
	if err := config.DB.Model(&book).Preload("Book_Category").Association("categories_id").Find(&categories); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Berhasil mendapatkan data buku dan kategori", "data": categories})
}

// BookGetImage godoc
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/book/image/{id} [get]
func BookGetImage(c *fiber.Ctx) error {
	idParam := c.Params("id")

	var book models.Books
	if err := config.DB.First(&book, idParam).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Terjadi kesalahan dalam mengambil data buku"})
	}

	pathImage := sampulPath + book.Img
	if _, err := os.Stat(pathImage); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Gambar tidak ditemukan"})
	}

	return c.SendFile(pathImage)
}
