package repository

import (
	"rest-api-redis/pkg/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *UserRepository) {
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
	userRepository := &UserRepository{DB: gdb}

	return gdb, mock, userRepository
}

func TestCreate(t *testing.T) {
	gdb, mock, userRepository := setupTest(t)
	defer gdb.Close()

	// Set up expectations for the query
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (.+) VALUES (.+)").
		WithArgs("John Doe", "john.doe@example.com", 25, sqlmock.AnyArg(), "123456", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the Create method
	user := &models.User{
		Name:         "John Doe",
		Email:        String("john.doe@example.com"),
		Age:          25,
		Birthday:     nil,
		MemberNumber: String("123456"),
		ActivatedAt:  nil,
	}

	createdUser, err := userRepository.Create(user)

	// Check for errors and expectations
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "John Doe", createdUser.Name)
	assert.NotEqual(t, 0, createdUser.ID) // Check that the ID is set

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindAll(t *testing.T) {
	gdb, mock, userRepository := setupTest(t)
	defer gdb.Close()

	// Set up expectations for the query
	mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "age", "birthday", "member_number", "activated_at", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john.doe@example.com", 25, nil, "123456", nil, time.Now(), time.Now()))

	// Call the FindAll method
	users, err := userRepository.FindAll(1, 10)

	// Check for errors and expectations
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, 1)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByID(t *testing.T) {
	gdb, mock, userRepository := setupTest(t)
	defer gdb.Close()

	// Set up expectations for the query
	mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "age", "birthday", "member_number", "activated_at", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john.doe@example.com", 25, nil, "123456", nil, time.Now(), time.Now()))

	// Call the FindByID method
	user, err := userRepository.FindByID(1)

	// Check for errors and expectations
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	gdb, mock, userRepository := setupTest(t)
	defer gdb.Close()

	// Set up expectations for the query
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `users` WHERE (.+)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the Delete method
	err := userRepository.Delete(1)

	// Check for errors and expectations
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	gdb, mock, userRepository := setupTest(t)
	defer gdb.Close()

	userToUpdate := &models.User{
		ID:           1,
		Name:         "Updated Name",
		Email:        String("updated.email@example.com"),
		Age:          30,
		Birthday:     nil,
		MemberNumber: String("654321"),
		ActivatedAt:  nil,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET (.+) WHERE `users`.`id` = ?").
		WithArgs("Updated Name", "updated.email@example.com", 30, nil, "654321", nil, time.Now(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	updatedUser, err := userRepository.Update(userToUpdate)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func String(s string) *string {
	return &s
}
