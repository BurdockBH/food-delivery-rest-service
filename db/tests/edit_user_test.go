package tests

import (
	"fmt"
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

func TestEditUser_Fail(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "edocicak@gmail.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL EditUser").ExpectQuery().WithArgs(
		u.Name, u.Email, sqlmock.AnyArg(), u.Phone, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"-1"}).AddRow(-1))

	err = user.EditUser(u)
	assert.Error(t, err)

	mock.ExpectPrepare("CALL EditUser").ExpectQuery().WithArgs(
		u.Name, u.Email, sqlmock.AnyArg(), u.Phone, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"-2"}).AddRow(-2))

	err = user.EditUser(u)
	assert.Error(t, err)
}

func TestEditUser_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL EditUser").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL EditUser").ExpectQuery().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
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
		err = user.EditUser(u)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}

}
