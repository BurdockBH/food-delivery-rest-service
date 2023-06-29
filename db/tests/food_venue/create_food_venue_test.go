package food_venue

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/food_venue"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateVenue_Success(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL CreateFoodVenue").ExpectExec().
		WithArgs(fv.Name, fv.Address, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))

	err = food_venue.CreateFoodVenue(&fv, "email@example.com")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateVenue_VenueExists(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL CreateFoodVenue").ExpectExec().
		WithArgs(fv.Name, fv.Address, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 0))

	err = food_venue.CreateFoodVenue(&fv, "email@example.com")
	assert.Error(t, err)
	assert.EqualError(t, err, "error with rows affected: <nil>")

}

func TestCreateVenue_ArgumentsError(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL CreateFoodVenue").ExpectExec().WithArgs(
		fv.Name, fv.Address, "email@example.com").WillReturnError(fmt.Errorf("Query 'CALL CreateFoodVenue(?, ?, ?)', arguments do not match: expected 3, but got 2 arguments"))

	err = food_venue.CreateFoodVenue(&fv, "email@example.com")
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL CreateFoodVenue(?, ?, ?)', arguments do not match: expected 3, but got 2 arguments")
}

func TestCreateVenue_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL CreateFoodVenue").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL CreateFoodVenue").ExpectExec().
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
		err = food_venue.CreateFoodVenue(&fv, "email@example.com")
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
