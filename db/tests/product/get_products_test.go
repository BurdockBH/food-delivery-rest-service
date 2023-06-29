package product

import (
	"errors"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProducts_Success(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	mock.ExpectPrepare("CALL GetProducts").ExpectQuery().WithArgs(
		sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{
		"id", "name", "description", "price", "quantity", "food_venue_id", "created_by", "created_at", "updated_at", "food_venue_name", "food_venue_address", "created_by", "created_at", "updated_at"}).
		AddRow(1, "name", "description", "price", 5, 3, "test@example.com", 1123123, 2131431, "venue", "address", "test@example.com", 1123123, 2131431))

	products, err := product.GetProducts(1)
	assert.NoError(t, err)
	assert.NotNil(t, products)
}

func TestGetProducts_Failure(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	mock.ExpectPrepare("CALL GetProducts").ExpectQuery().WithArgs(
		sqlmock.AnyArg()).WillReturnError(errors.New("database error"))

	products, err := product.GetProducts(1)
	assert.Error(t, err)
	assert.Nil(t, products)

}

func TestGetProducts_ArgumentsError(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	mock.ExpectPrepare("CALL GetProducts").ExpectQuery().WithArgs(
		sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL GetProducts(?)', arguments do not match: expected 2, but got 1 arguments"))

	_, err = product.GetProducts(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL GetProducts(?)', arguments do not match: expected 2, but got 1 arguments")
}

func TestGetProducts_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL GetProducts").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL GetProducts").ExpectQuery().
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	for _, data := range testData {
		data.mockFn(data.err)
		_, err = product.GetProducts(1)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
