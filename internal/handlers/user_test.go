package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// User represents a user entity.
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	Birthday     time.Time `json:"birthday"`
	MemberNumber string    `json:"member_number"`
	ActivatedAt  time.Time `json:"activated_at"`
}

// CreateUserPayload represents the payload for creating a user.
type CreateUserPayload struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          any       `json:"age"`
	Birthday     time.Time `json:"birthday"`
	MemberNumber string    `json:"member_number"`
	ActivatedAt  time.Time `json:"activated_at"`
}

// UpdateUserPayload represents the payload for updating a user.
type UpdateUserPayload struct {
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Age          int        `json:"age"`
	Birthday     *time.Time `json:"birthday"`
	MemberNumber string     `json:"member_number"`
	ActivatedAt  *time.Time `json:"activated_at"`
}

// Database represents the database operations.
type Database interface {
	CreateUser(user *CreateUserPayload) error
	GetUser(userID int) (*User, error)
	GetUsers(page, limit int) ([]*User, error)
	DeleteUser(userID int) error
	UpdateUser(userID int, user *UpdateUserPayload) (*User, error)
}

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// type CreateUserPayload struct {
// 	Name         string    `json:"name"`
// 	Email        string    `json:"email"`
// 	Age          any       `json:"age"`
// 	Birthday     time.Time `json:"birthday"`
// 	MemberNumber string    `json:"member_number"`
// 	ActivatedAt  time.Time `json:"activated_at"`
// }

func TestCreateUser(t *testing.T) {
	app := fiber.New()
	app.Post("/users", CreateUser)

	tests := []struct {
		name           string
		payload        CreateUserPayload
		expectedStatus int
	}{
		{
			name: "ValidPayload",
			payload: CreateUserPayload{
				Name:         "Nidhey",
				Email:        "nidhey60@gmail.com",
				Age:          25,
				Birthday:     time.Date(1999, 3, 27, 0, 0, 0, 0, time.UTC),
				MemberNumber: "27031999",
				ActivatedAt:  time.Date(2023, 1, 1, 12, 30, 45, 0, time.UTC),
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Payload",
			payload: CreateUserPayload{
				Email:        "nidhey60@gmail.com",
				Age:          25,
				Birthday:     time.Date(1999, 3, 27, 0, 0, 0, 0, time.UTC),
				MemberNumber: "27031999",
				ActivatedAt:  time.Date(2023, 1, 1, 12, 30, 45, 0, time.UTC),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "MalformedPayload",
			payload: CreateUserPayload{
				Name:         "Nidhey Indurkar",
				Email:        "nidhey60@gmail.com",
				Age:          "invalid",
				Birthday:     time.Date(1999, 3, 27, 0, 0, 0, 0, time.UTC),
				MemberNumber: "27031999",
				ActivatedAt:  time.Date(2023, 1, 1, 12, 30, 45, 0, time.UTC),
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "RepositoryError",
			payload: CreateUserPayload{
				Name:         "Nidhey Indurkar",
				Email:        "nidhey60@gmail.com",
				Age:          "invalid",
				Birthday:     time.Date(1999, 3, 27, 0, 0, 0, 0, time.UTC),
				MemberNumber: "27031999",
				ActivatedAt:  time.Date(2023, 1, 1, 12, 30, 45, 0, time.UTC),
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payload, err := json.Marshal(test.payload)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(payload)))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatus, resp.StatusCode)
		})
	}
}

// func TestGetUser(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/users/:id", GetUser)

// 	t.Run("GetUser", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodGet, "/users/39", nil)
// 		resp, err := app.Test(req)

// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
// }

// func TestGetUsers(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/users", GetUsers)

// 	t.Run("GetUsers", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodGet, "/users?page=1&limit=10", nil)
// 		resp, err := app.Test(req)

// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
// }

// func TestDeleteUser(t *testing.T) {
// 	app := fiber.New()
// 	app.Delete("/users/:id", DeleteUser)

// 	t.Run("DeleteUser", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodDelete, "/users/40", nil)
// 		resp, err := app.Test(req)

// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
// }

// func TestUpdateUser(t *testing.T) {
// 	app := fiber.New()
// 	app.Put("/users/:id", UpdateUser)

// 	t.Run("UpdateUser", func(t *testing.T) {
// 		payload := `{"name": "Updated Name", "email": "updated.email@example.com", "age": 30, "birthday": null, "member_number": "654321", "activated_at": null}`
// 		req := httptest.NewRequest(http.MethodPut, "/users/41", strings.NewReader(payload))
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, err := app.Test(req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	})
// }

type DatabaseMock struct {
	mock.Mock
}

func (m *DatabaseMock) CreateUser(user *CreateUserPayload) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *DatabaseMock) GetUser(userID int) (*User, error) {
	args := m.Called(userID)
	return args.Get(0).(*User), args.Error(1)
}

func (m *DatabaseMock) GetUsers(page, limit int) ([]*User, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]*User), args.Error(1)
}

func (m *DatabaseMock) DeleteUser(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *DatabaseMock) UpdateUser(userID int, user *UpdateUserPayload) (*User, error) {
	args := m.Called(userID, user)
	return args.Get(0).(*User), args.Error(1)
}
