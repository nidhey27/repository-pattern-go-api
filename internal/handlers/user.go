package handlers

import (
	"fmt"
	"net/http"
	"rest-api-redis/pkg/database"
	"rest-api-redis/pkg/models"
	"rest-api-redis/pkg/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := &models.User{}
	db := database.GetDB()

	if err := c.BodyParser(&user); err != nil {
		return utils.SendResponse(http.StatusBadRequest, "", err.Error(), make([]string, 0), c)
	}

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	result := db.Create(user)

	if result.Error != nil {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusCreated, "User created successfully", "", user, c)
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}
	user := &models.User{}
	db := database.GetDB()

	result := db.First(&user, "id = ?", userId)

	if result.Error != nil {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, "User fetched successfully", "", user, c)
}

func GetUsers(c *fiber.Ctx) error {

	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	users := []models.User{}
	db := database.GetDB()

	result := db.Limit(intLimit).Offset(offset).Find(&users)

	if result.Error != nil {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}
	responseData := struct {
		Users      []models.User `json:"users"`
		Page       int           `json:"page"`
		Limit      int           `json:"limit"`
		TotalCount int64         `json:"total_count"`
	}{
		Users:      users,
		Page:       intPage,
		Limit:      intLimit,
		TotalCount: result.RowsAffected,
	}

	return utils.SendResponse(http.StatusOK, "Users fetched successfully", "", responseData, c)
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}
	user := &models.User{}

	db := database.GetDB()

	result := db.First(&user, "id = ?", userId)

	if result.Error != nil && strings.Contains(result.Error.Error(), "record not found") {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}

	result = db.Delete(&user, "id = ?", userId)

	if result.Error != nil {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, fmt.Sprintf("User with ID %v deleted successfully", userId), "", make([]string, 0), c)
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}
	user := &models.User{}

	db := database.GetDB()
	db.LogMode(true)

	result := db.First(&user, "id = ?", userId)

	if result.Error != nil && strings.Contains(result.Error.Error(), "record not found") {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}
	c.BodyParser(&user)

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	uid, _ := strconv.Atoi(userId)
	user.ID = uint(uid)
	result = db.Save(&user)

	if result.Error != nil {
		return utils.SendResponse(http.StatusBadGateway, "", result.Error.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, fmt.Sprintf("User with ID %v updated successfully", userId), "", result.RowsAffected, c)
}
