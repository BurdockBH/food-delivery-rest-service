package product

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProduct_Success(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	p := viewmodels.Product{
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
		FoodVenue: viewmodels.FoodVenue{
			Name:    "Venue",
			Address: "Address",
		},
	}

	mock.ExpectPrepare("CALL CreateProduct").ExpectQuery().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(
		sqlmock.NewRows([]string{"1"}).AddRow(1))

	err = product.CreateProduct(&p, "email@example.com")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateProduct_ProductExists(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	p := viewmodels.Product{
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
		FoodVenue: viewmodels.FoodVenue{
			Name:    "Venue",
			Address: "Address",
		},
	}

	mock.ExpectPrepare("CALL CreateProduct").ExpectQuery().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(
		sqlmock.NewRows([]string{"0"}).AddRow(0))

	err = product.CreateProduct(&p, "email@example.com")
	assert.Error(t, err)

}

func TestCreateProduct_ArgumentsError(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	p := viewmodels.Product{
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
		FoodVenue: viewmodels.FoodVenue{
			Name:    "Venue",
			Address: "Address",
		},
	}

	mock.ExpectPrepare("CALL CreateProduct").ExpectQuery().WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL CreateProduct(?, ?, ?, ?, ?, ? ,?)', arguments do not match: expected 7, but got 6 arguments"))

	err = product.CreateProduct(&p, "email@example.com")
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL CreateProduct(?, ?, ?, ?, ?, ? ,?)', arguments do not match: expected 7, but got 6 arguments")
}

func TestCreateProduct_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL CreateProduct").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL CreateProduct").ExpectQuery().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	p := viewmodels.Product{
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
		FoodVenue: viewmodels.FoodVenue{
			Name:    "Venue",
			Address: "Address",
		},
	}

	for _, data := range testData {
		data.mockFn(data.err)
		err = product.CreateProduct(&p, "email@example.com")
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
