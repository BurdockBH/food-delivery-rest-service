package product

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteProduct_Success(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	mock.ExpectPrepare("CALL DeleteProduct").ExpectQuery().
		WithArgs(sqlmock.AnyArg()).WillReturnRows(
		sqlmock.NewRows([]string{"1"}).AddRow(1))

	err = product.DeleteProduct(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteProduct_Failure(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	mock.ExpectPrepare("CALL DeleteProduct").ExpectQuery().
		WithArgs(sqlmock.AnyArg()).WillReturnRows(
		sqlmock.NewRows([]string{"0"}).AddRow(0))

	err = product.DeleteProduct(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Product with id 1 does not exist")

}

func TestDeleteProduct_ArgumentsError(t *testing.T) {
	db2, mock, err := helper.MockDatabase()
	assert.NoError(t, err)
	defer db2.Close()

	mock.ExpectPrepare("CALL DeleteProduct").ExpectQuery().WithArgs(
		sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL DeleteProduct(?)', arguments do not match: expected 1, but got 2 arguments"))

	err = product.DeleteProduct(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL DeleteProduct(?)', arguments do not match: expected 1, but got 2 arguments")
}

func TestDeleteProduct_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL DeleteProduct").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL DeleteProduct").ExpectQuery().
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	for _, data := range testData {
		data.mockFn(data.err)
		err = product.DeleteProduct(1)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
