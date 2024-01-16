package handlers

import (
	"net/http"
	"net/http/httptest"
	"rest-api-redis/pkg/repository"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (sqlmock.Sqlmock, *repository.UserRepository) {
	// Set up a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}

	// Initialize GORM with the mock DB
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("Error creating GORM instance: %v", err)
	}

	// Create UserRepository instance with the mock DB
	userRepository := &repository.UserRepository{DB: gdb}

	return mock, userRepository
}

func TestCreateUser(t *testing.T) {
	scenarios := []struct {
		name       string
		userJSON   string
		statusCode int
	}{
		{"All Fields Provided", `{"name": "John Doe", "email": "john.doe@example.com", "age": 25}`, http.StatusCreated},
		{"Name and Email Provided", `{"name": "John Doe", "email": "john.doe@example.com"}`, http.StatusBadRequest},
		{"Email and Age Provided", `{"email": "john.doe@example.com", "age": 25}`, http.StatusBadRequest},
		{"Age and Name Provided", `{"name": "John Doe", "age": 25}`, http.StatusBadRequest},
		{"No Body", ``, http.StatusBadRequest},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			mock, userRepository := setupTest(t)

			// Mock repository
			handler := InitUserHandler(userRepository)

			// Create a request and response recorder
			req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(scenario.userJSON))
			req.Header.Set("Content-Type", "application/json")

			// Call the CreateUser handler
			app := fiber.New()
			app.Post("/user", handler.CreateUser)

			// Set up expectations for the mock database
			if scenario.statusCode == http.StatusCreated {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users` (.+)").WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			resp, err := app.Test(req)
			assert.NoError(t, err)

			// Assert the response status code
			assert.Equal(t, scenario.statusCode, resp.StatusCode)
		})
	}
}

func TestGetUser(t *testing.T) {
	mock, userRepository := setupTest(t)

	// Mock repository
	handler := InitUserHandler(userRepository)

	mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "age", "birthday", "member_number", "activated_at", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john.doe@example.com", 25, nil, "123456", nil, time.Now(), time.Now()))

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/user/1", nil)
	req.Header.Set("Content-Type", "application/json")

	// Call the GetUser handler
	app := fiber.New()
	app.Get("/user/:id", handler.GetUser)

	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetUsers(t *testing.T) {
	mock, userRepository := setupTest(t)

	// Mock repository
	handler := InitUserHandler(userRepository)

	// Mock data
	mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "age", "birthday", "member_number", "activated_at", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john.doe@example.com", 25, nil, "123456", nil, time.Now(), time.Now()))

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	// Call the GetUsers handler
	app := fiber.New()
	app.Get("/users", handler.GetUsers)

	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteUser(t *testing.T) {
	mock, userRepository := setupTest(t)

	// Mock repository
	handler := InitUserHandler(userRepository)

	// Mock data
	mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "age", "birthday", "member_number", "activated_at", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john.doe@example.com", 25, nil, "123456", nil, time.Now(), time.Now()))
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `users` WHERE (.+)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	req.Header.Set("Content-Type", "application/json")

	// Call the DeleteUser handler
	app := fiber.New()
	app.Delete("/user/:id", handler.DeleteUser)

	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateUser(t *testing.T) {
	mock, userRepository := setupTest(t)

	// Mock repository
	handler := InitUserHandler(userRepository)

	// Mock data
	mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "age", "birthday", "member_number"}).
			AddRow(1, "John Doe", "john.doe@example.com", 25, nil, "123456"))
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET (.+) WHERE `users`.`id` = ?").
		WithArgs("Updated Name", "updated.email@example.com", 26, nil, "123456", nil, sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodPut, "/user/1", strings.NewReader(`{"name": "Updated Name", "email": "updated.email@example.com", "age": 26, "member_number": "123456"}`))
	req.Header.Set("Content-Type", "application/json")

	// Call the UpdateUser handler
	app := fiber.New()
	app.Put("/user/:id", handler.UpdateUser)

	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
