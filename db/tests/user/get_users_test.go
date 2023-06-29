package user

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUsers_Success(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL GetUsersByDetails").ExpectQuery().WithArgs(
		u.Name, u.Email, u.Phone).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	_, err = user.GetUsers(u)
	assert.NoError(t, err)
}

func TestGetUsers_Fail(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL GetUsersByDetails").ExpectQuery().WithArgs(
		u.Name, u.Email, u.Phone).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name", "password", "phone", "created_at", "updated_at"}))

	users, err := user.GetUsers(u)
	assert.Error(t, err)
	assert.Nil(t, users)
}

func TestGetUsers_ArgumentError(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL GetUsersByDetails").ExpectQuery().WithArgs(
		u.Name, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL GetUsersByDetails(?, ?, ?)', arguments do not match: expected 3, but got 2 arguments"))

	_, err = user.GetUsers(u)
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL GetUsersByDetails(?, ?, ?)', arguments do not match: expected 3, but got 2 arguments")
}

func TestGetUsers_PrepareExec(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	testData := []struct {
		err    error
		mockFn func(err error)
	}{
		{
			err: fmt.Errorf("preparation error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL GetUsers").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL GetUsers").ExpectQuery().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	for _, data := range testData {
		data.mockFn(data.err)
		_, err = user.GetUsers(u)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
