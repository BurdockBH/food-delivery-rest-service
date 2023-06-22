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

func TestLoginUser(t *testing.T) {
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
	loginUser := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	hashedPassword, err := helper.HashPassword(registerUser.Password)
	assert.NoError(t, err)
	//
	//mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().WithArgs(
	//	registerUser.Name, registerUser.Email, sqlmock.AnyArg(), registerUser.Phone, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	//
	//err = user.RegisterUser(registerUser)
	//assert.NoError(t, err)

	mock.ExpectPrepare("CALL LoginUser").ExpectQuery().WithArgs(
		loginUser.Email).WillReturnRows(sqlmock.NewRows([]string{"password"})).WithArgs(hashedPassword)

	err = user.LoginUser(loginUser)
	assert.NoError(t, err)
}
