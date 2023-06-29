package user

import (
	"fmt"
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

	loginUser := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	mock.ExpectPrepare("CALL LoginUser").ExpectQuery().WithArgs(
		loginUser.Email).WillReturnRows(sqlmock.NewRows([]string{"0"}))

	err = user.LoginUser(loginUser)
	assert.EqualError(t, err, "user with email edocicak@gmail.com does not exist")
}

func TestLoginUser_ArgumentsError(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	u := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	mock.ExpectPrepare("CALL LoginUser").ExpectQuery().WithArgs(
		u.Email).WillReturnError(fmt.Errorf("Query 'CALL LoginUser(?)', arguments do not match: expected 1, but got 2 arguments"))

	err = user.LoginUser(u)

	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL LoginUser(?)', arguments do not match: expected 1, but got 2 arguments")
}

func TestLoginUser_PrepareExec(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	testData := []struct {
		err    error
		mockFn func(err error)
	}{
		{
			err: fmt.Errorf("preparation error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL LoginUser").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL LoginUser").ExpectQuery().
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	u := viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	for _, data := range testData {
		data.mockFn(data.err)
		err = user.LoginUser(&u)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
