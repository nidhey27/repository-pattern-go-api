package handlers

import (
	"fmt"
	"net/http"
	"rest-api-redis/pkg/database"
	"rest-api-redis/pkg/models"
	"rest-api-redis/pkg/repository"
	"rest-api-redis/pkg/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := &models.User{}
	db := database.GetDB()
	userRepository := repository.ProvideUserRepository(db)

	if err := c.BodyParser(&user); err != nil {
		return utils.SendResponse(http.StatusBadRequest, "", err.Error(), make([]string, 0), c)
	}

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	_, err := userRepository.Create(user)

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusCreated, "User created successfully", "", user, c)
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}
	db := database.GetDB()
	userRepository := repository.ProvideUserRepository(db)

	id, _ := strconv.Atoi(userId)
	result, err := userRepository.FindByID(uint(id))

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, "User fetched successfully", "", result, c)
}

func GetUsers(c *fiber.Ctx) error {

	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	db := database.GetDB()
	userRepository := repository.ProvideUserRepository(db)

	result, err := userRepository.FindAll(intLimit, offset)

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}
	responseData := struct {
		Users      []models.User `json:"users"`
		Page       int           `json:"page"`
		Limit      int           `json:"limit"`
		TotalCount int           `json:"total_count"`
	}{
		Users:      result,
		Page:       intPage,
		Limit:      intLimit,
		TotalCount: len(result),
	}

	return utils.SendResponse(http.StatusOK, "Users fetched successfully", "", responseData, c)
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}

	db := database.GetDB()
	userRepository := repository.ProvideUserRepository(db)

	id, _ := strconv.Atoi(userId)
	_, err := userRepository.FindByID(uint(id))

	if err != nil && strings.Contains(err.Error(), "record not found") {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	err = userRepository.Delete(uint(id))

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
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
	userRepository := repository.ProvideUserRepository(db)

	id, _ := strconv.Atoi(userId)
	_, err := userRepository.FindByID(uint(id))

	if err != nil && strings.Contains(err.Error(), "record not found") {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}
	c.BodyParser(&user)

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	uid, _ := strconv.Atoi(userId)
	user.ID = uint(uid)
	updatedUser, err := userRepository.Update(user)

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, fmt.Sprintf("User with ID %v updated successfully", userId), "", updatedUser, c)
}
