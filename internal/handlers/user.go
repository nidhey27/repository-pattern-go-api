package handlers

import (
	"fmt"
	"net/http"
	"rest-api-redis/pkg/models"
	"rest-api-redis/pkg/repository"
	"rest-api-redis/pkg/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserHandler struct {
	service repository.IUserRepository
}

func InitUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		service: userRepo,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := &models.User{}
	if err := c.BodyParser(&user); err != nil {
		return utils.SendResponse(http.StatusBadRequest, "", err.Error(), make([]string, 0), c)
	}

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	_, err := h.service.Create(user)

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusCreated, "User created successfully", "", user, c)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}

	id, _ := strconv.Atoi(userId)
	result, err := h.service.FindByID(uint(id))

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, "User fetched successfully", "", result, c)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {

	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	result, err := h.service.FindAll(intLimit, offset)

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

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}

	id, _ := strconv.Atoi(userId)
	_, err := h.service.FindByID(uint(id))

	if err != nil && strings.Contains(err.Error(), "record not found") {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	err = h.service.Delete(uint(id))

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, fmt.Sprintf("User with ID %v deleted successfully", userId), "", make([]string, 0), c)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if userId == "" {
		return utils.SendResponse(http.StatusBadRequest, "", "ID is required", make([]string, 0), c)
	}
	user := &models.User{}

	id, _ := strconv.Atoi(userId)
	_, err := h.service.FindByID(uint(id))

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
	updatedUser, err := h.service.Update(user)

	if err != nil {
		return utils.SendResponse(http.StatusBadGateway, "", err.Error(), make([]string, 0), c)
	}

	return utils.SendResponse(http.StatusOK, fmt.Sprintf("User with ID %v updated successfully", userId), "", updatedUser, c)
}
