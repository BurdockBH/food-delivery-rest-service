package tests

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

func TestDeleteUse_Success(t *testing.T) {
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

	mockRows := sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword)
	mockRows2 := sqlmock.NewRows([]string{"1"}).AddRow(1)

	mock.ExpectQuery("CALL LoginUser").WithArgs(
		loginUser.Email).WillReturnRows(mockRows)

	mock.ExpectPrepare("CALL DeleteUser").ExpectQuery().WithArgs(
		loginUser.Email).WillReturnRows(mockRows2)

	err = user.DeleteUser(loginUser)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser_NoUser(t *testing.T) {
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

	mockRows := sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword)
	mockRows2 := sqlmock.NewRows([]string{"0"}).AddRow(0)

	mock.ExpectQuery("CALL LoginUser").WithArgs(
		loginUser.Email).WillReturnRows(mockRows)

	mock.ExpectPrepare("CALL DeleteUser").ExpectQuery().WithArgs(
		loginUser.Email).WillReturnRows(mockRows2)

	err = user.DeleteUser(loginUser)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser_ArgumentsError(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	u := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	hashedPassword, err := helper.HashPassword(u.Password)
	assert.NoError(t, err)

	mockRows := sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword)

	mock.ExpectQuery("CALL LoginUser").WithArgs(
		u.Email).WillReturnRows(mockRows)

	mock.ExpectPrepare("CALL DeleteUser").ExpectQuery().WithArgs(
		u.Email).WillReturnError(fmt.Errorf("Query 'CALL DeleteUser(?)', arguments do not match: expected 1, but got 2 arguments"))

	err = user.DeleteUser(u)

	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL DeleteUser(?)', arguments do not match: expected 1, but got 2 arguments")
}

func TestDeleteUser_PrepareExec(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	u := &viewmodels.UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	hashedPassword, err := helper.HashPassword(u.Password)
	assert.NoError(t, err)

	testData := []struct {
		err    error
		mockFn func(err error)
	}{
		{
			err: fmt.Errorf("preparation error"),
			mockFn: func(err error) {
				mock.ExpectQuery("CALL LoginUser").WithArgs(
					u.Email).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword))
				mock.ExpectPrepare("CALL DeleteUser").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectQuery("CALL LoginUser").WithArgs(
					u.Email).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword))
				mock.ExpectPrepare("CALL DeleteUser").ExpectQuery().
					WithArgs(u.Email).
					WillReturnError(err)
			},
		},
	}

	for _, data := range testData {
		data.mockFn(data.err)
		err = user.DeleteUser(u)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
