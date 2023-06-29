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

func TestDeleteVenue_Success(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL DeleteFoodVenue").ExpectQuery().
		WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	err = food_venue.DeleteFoodVenue(&fv)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteVenue_Failure(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL DeleteFoodVenue").ExpectQuery().
		WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"0"}).AddRow(0))

	err = food_venue.DeleteFoodVenue(&fv)
	assert.Error(t, err)
	assert.EqualError(t, err, "Food venue with id 0 does not exist")

}

func TestDeleteVenue_ArgumentsError(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	fv := viewmodels.FoodVenue{
		Name:    "Venue",
		Address: "Address",
	}

	mock.ExpectPrepare("CALL DeleteFoodVenue").ExpectQuery().WithArgs(
		sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL DeleteFoodVenue(?)', arguments do not match: expected 1, but got 2 arguments"))

	err = food_venue.DeleteFoodVenue(&fv)
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL DeleteFoodVenue(?)', arguments do not match: expected 1, but got 2 arguments")
}

func TestDeleteVenue_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL DeleteFoodVenue").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL DeleteFoodVenue").ExpectQuery().
					WithArgs(sqlmock.AnyArg()).
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
		err = food_venue.DeleteFoodVenue(&fv)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
