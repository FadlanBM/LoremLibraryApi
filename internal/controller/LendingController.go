package controller

import (
	"errors"
	"github.com/api/config"
	helper "github.com/api/internal/helpers"
	"github.com/api/internal/models"
	"github.com/api/internal/request"
	response "github.com/api/internal/respose"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// LendingIndex godoc
// @Tags Lendings
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/lending [get]
func LendingIndex(c *fiber.Ctx) error {
	var lending []models.Lending
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

	if err := config.DB.Find(&lending, "borrowers_id= ?", borrowers.ID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var res []response.LendingResponse

	for _, v := range lending {
		var dateLast string
		dateLast = ""
		if v.DateLast != nil {
			dateLast = helper.FormatDate(*v.DateLast)
		}

		transaksiResponse := response.LendingResponse{
			ID:         v.ID,
			DateLast:   dateLast,
			ReturnDate: helper.FormatDate(v.ReturnDate),
			BorrowDate: helper.FormatDate(v.BorrowDate),
			Code:       v.Code,
			Status:     v.Status,
		}
		res = append(res, transaksiResponse)
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Succcess get data", "data": res})
}

// LendingsCreate godoc
// @Tags Lendings
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Param request body request.LendingRequest true "Request"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/lending/create [post]
func LendingsCreate(c *fiber.Ctx) error {
	req := new(request.LendingRequest)
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
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data not found"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	tglterbit, err := time.Parse("02/01/2006", req.ReturnDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Invalid date format"})
	}

	lendings := models.Lending{
		ReturnDate:  tglterbit,
		BorrowDate:  time.Now(),
		BorrowersID: borrowers.ID,
		DateLast:    nil,
		Code:        helper.GenerateRandomString(8),
	}

	if err := config.DB.Create(&lendings).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Berhasil melakukan transaksi", "data": lendings})
}

// ListLendingCreate godoc
// @Tags Lendings
// @Accept json
// @Produce json
// @Param request body request.ListLending true "Request"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/lending/list [post]
func ListLendingCreate(c *fiber.Ctx) error {
	req := new(request.ListLending)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	listlending := models.ListLending{
		BooksID:       req.BookID,
		LendingID:     req.LendingID,
		No_Inventaris: req.NoInventaris,
	}

	if err := config.DB.Create(&listlending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Berhasil melakukan transaksi"})
}

// DetailLending godoc
// @Tags Lendings
// @Accept json
// @Produce json
// @Param id path int true "Lendings ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Router /api/lending/{id} [get]
func DetailLending(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var lendings models.Lending
	if err := config.DB.First(&lendings, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Success ambil data ", "data": lendings})
}

// HistroyLending godoc
// @Tags Lendings
// @Summary Authenticate a user and generate a token
// @Description Authenticate a user and generate a token based on the provided credentials.
// @Accept json
// @Produce json
// @Param id path int true "Lendings ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Failure 404 {object} response.ResponseError
// @Router /api/lending/history/{id} [get]
func HistroyLending(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var lending models.Lending
	if err := config.DB.First(&lending, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data tidak di temukan"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var listLending []models.ListLending

	if err := config.DB.Find(&listLending, "lending_id = ?", lending.ID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	var listlendingID []int
	for _, bc := range listLending {
		listlendingID = append(listlendingID, int(bc.BooksID))
	}

	var book []models.Books
	if err := config.DB.Where("id IN (?)", listlendingID).Find(&book).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": err.Error()})
	}

	bookWithCategories := struct {
		Book []models.Books
	}{
		Book: book,
	}

	return c.Status(200).JSON(bookWithCategories)
}

// DeleteLending godoc
// @Tags Lendings
// @Accept json
// @Produce json
// @Param id path int true "Lendings ID"
// @Success 200 {object} response.ResponseDataSuccess
// @Failure 400 {object} response.ResponseError
// @Failure 404 {object} response.ResponseError
// @Router /api/lending/delete/{id} [delete]
func DeleteLending(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "false", "message": "Invalid ID format"})
	}

	var lending models.Lending
	if err := config.DB.First(&lending, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "false", "message": "Data Not Found"})
		}
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to retrieve lending data"})
	}

	var listLending []models.ListLending
	if err := config.DB.Find(&listLending, "lending_id = ?", lending.ID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to retrieve list lending data"})
	}

	if len(listLending) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "false", "message": "List lending data not found"})
	}

	if res := config.DB.Unscoped().Delete(&listLending); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete list lending data"})
	}

	if res := config.DB.Unscoped().Delete(&lending); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "false", "message": "Failed to delete lending data"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "true", "message": "Success delete data"})

}
