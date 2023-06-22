package tests

import (
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterUser_Success(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2
	defer func() { db2 = db.DB }()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().
		WithArgs(u.Name, u.Email, sqlmock.AnyArg(), u.Phone, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"created"}).AddRow(1))

	err = user.RegisterUser(u)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegisterUserExists_Success(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2
	defer func() { db2 = db.DB }()

	u := &viewmodels.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().
		WithArgs(u.Name, u.Email, sqlmock.AnyArg(), u.Phone, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	err = user.RegisterUser(u)
	assert.NoError(t, err)

	mock.ExpectPrepare("CALL RegisterUser").ExpectQuery().
		WithArgs(u.Name, u.Email, sqlmock.AnyArg(), u.Phone, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"0"}))

	err = user.RegisterUser(u)

	assert.Error(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")

	assert.NoError(t, mock.ExpectationsWereMet())
}
