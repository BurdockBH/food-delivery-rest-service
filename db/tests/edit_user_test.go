package tests

import (
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditUser(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2
	defer func() { db2 = db.DB }()

	registerUser := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL EditUser").ExpectQuery().WithArgs(
		registerUser.Name, registerUser.Email, sqlmock.AnyArg(), registerUser.Phone, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Gives me user edocicak@gmail does not exist
	err = user.EditUser(registerUser)
	assert.NoError(t, err)
}
