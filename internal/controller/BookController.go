package controller

import (
	"errors"
	"github.com/api/config"
	"github.com/api/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
)

const (
	sampulPath = "E:\\UKK\\LoremLibrary\\public\\storage\\sampul\\"
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
// @Param id path int true "petugas ID"
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
// @Router /api/book/category/{id} [get]
func CategoryDetail(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "ID harus berupa angka"})
	}

	var books models.Books
	if err := config.DB.First(&books, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var bookCategory []models.Book_Category

	if err := config.DB.Find(&bookCategory, "books_id = ?", books.ID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var categoryIDs []int
	for _, bc := range bookCategory {
		categoryIDs = append(categoryIDs, int(bc.CategoriesID))
	}

	var categories []models.Categories
	if err := config.DB.Where("id IN (?)", categoryIDs).Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	// Menggabungkan data buku dan kategori
	bookWithCategories := struct {
		Categories []models.Categories
	}{
		Categories: categories,
	}

	return c.Status(200).JSON(bookWithCategories)
}

// BookGetImage godoc
// @Tags Book
// @Accept json
// @Produce json
// @Param path path string true "path"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/book/image/{path} [get]
func BookGetImage(c *fiber.Ctx) error {
	path := c.Params("path")

	// Pastikan path gambar valid
	if filepath.Ext(path) == "" {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Path tidak valid"})
	}

	pathImage := filepath.Join(sampulPath, path)
	if _, err := os.Stat(pathImage); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Gambar tidak ditemukan"})
	}

	return c.SendFile(pathImage)
}
