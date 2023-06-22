package tests

import (
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2
	defer func() { db2 = db.DB }()

	loginUser := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	// Gives me user edocicak@gmail.com does not exist
	mock.ExpectPrepare("CALL DeleteUser").ExpectExec().WithArgs(
		loginUser.Email).WillReturnResult(sqlmock.NewResult(0, 1))

	err = user.DeleteUser(loginUser)
	assert.NoError(t, err)
}
