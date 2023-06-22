package tests

import (
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginUser_Success(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2
	defer func() { db2 = db.DB }()

	loginUser := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	hashedPassword, err := helper.HashPassword(loginUser.Password)
	assert.NoError(t, err)

	mock.ExpectPrepare("CALL LoginUser").ExpectQuery().WithArgs(
		loginUser.Email).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword))

	err = user.LoginUser(loginUser)
	assert.NoError(t, err)
}

func TestLoginUser_NoUser(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2
	defer func() { db2 = db.DB }()

	loginUser := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	mock.ExpectPrepare("CALL LoginUser").ExpectQuery().WithArgs(
		loginUser.Email).WillReturnRows(sqlmock.NewRows([]string{"0"}))

	err = user.LoginUser(loginUser)
	assert.EqualError(t, err, "user edocicak@gmail.com does not exist")
}
