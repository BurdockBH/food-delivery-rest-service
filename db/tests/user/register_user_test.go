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

func TestRegisterUser_Success(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().
		WithArgs(u.Name, u.Email, sqlmock.AnyArg(), u.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	err = user.RegisterUser(u)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegisterUser_UserExists(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().
		WithArgs(u.Name, u.Email, sqlmock.AnyArg(), u.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"0"}).AddRow(0))

	err = user.RegisterUser(u)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("user with email %v or phone number %v already exists", u.Email, u.Phone))
}

func TestRegisterUser_ArgumentsError(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().WithArgs(
		u.Name, sqlmock.AnyArg(), u.Phone, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL RegisterUser(?, ?, ?, ?)', arguments do not match: expected 5, but got 4 arguments"))

	err = user.RegisterUser(u)
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL RegisterUser(?, ?, ?, ?)', arguments do not match: expected 5, but got 4 arguments")
}

func TestRegisterUser_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL RegisterUser").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
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
		err = user.RegisterUser(u)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
