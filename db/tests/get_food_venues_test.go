package tests

import (
	"errors"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/food_venue"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFoodVenues_Success(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL GetVenues").ExpectQuery().WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address", "created_by", "created_at", "updated_at"}).AddRow(
		1, "name", "address", "test@test.com", 1231452, 1123123))

	venues, err := food_venue.GetVenues(&fv)
	assert.NoError(t, err)
	assert.NotNil(t, venues)
}

func TestGetFoodVenues_Failure(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL GetVenues").ExpectQuery().WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("database error"))

	venues, err := food_venue.GetVenues(&fv)
	assert.Error(t, err)
	assert.Nil(t, venues)
}

func TestGetFoodVenues_ArgumentError(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL GetVenues").ExpectQuery().WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL GetVenues(?, ?, ?)', arguments do not match: expected 3, but got 2 arguments"))

	_, err = food_venue.GetVenues(&fv)

	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL GetVenues(?, ?, ?)', arguments do not match: expected 3, but got 2 arguments")
}

func TestGetFoodVenues_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL GetVenues").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL GetVenues").ExpectQuery().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	for _, data := range testData {
		data.mockFn(data.err)
		_, err = food_venue.GetVenues(&fv)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
