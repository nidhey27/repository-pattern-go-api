package handlers

// import (
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/require"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func TestProductRepository_FindByID(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	gdb, err := gorm.Open("mysql", db)
// 	productRepo := ProvideProductRepostiory(gdb)

// 	mock.ExpectQuery(
// 		"SELECT(.*)").
// 		WithArgs(1).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "code", "price"}).
// 				AddRow(products[0].ID, products[0].CreatedAt, products[0].UpdatedAt, products[0].DeletedAt, products[0].Code, products[0].Price))
// 	res, err := productRepo.FindByID(products[0].ID)

// 	require.NoError(t, err)
// 	require.Equal(t, res, products[0])
// }