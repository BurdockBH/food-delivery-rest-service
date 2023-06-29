package product

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditProduct_Success(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	p := viewmodels.Product{
		ID:          1,
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
	}

	mock.ExpectPrepare("CALL EditProduct").ExpectQuery().
		WithArgs(p.ID, p.Name, p.Description, p.Price, p.Quantity).WillReturnRows(
		sqlmock.NewRows([]string{"1"}).AddRow(1))

	err = product.EditProduct(&p)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEditProduct_Failure(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	p := viewmodels.Product{
		ID:          1,
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
	}

	mock.ExpectPrepare("CALL EditProduct").ExpectQuery().
		WithArgs(p.ID, p.Name, p.Description, p.Price, p.Quantity).WillReturnRows(
		sqlmock.NewRows([]string{"0"}).AddRow(0))

	err = product.EditProduct(&p)
	assert.Error(t, err)
	assert.EqualError(t, err, "Product with id 1 does not exist")

}

func TestEditProduct_ArgumentsError(t *testing.T) {
	db2, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db2.Close()

	db.DB = db2

	p := viewmodels.Product{
		ID:          1,
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
	}

	mock.ExpectPrepare("CALL EditProduct").ExpectQuery().WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(fmt.Errorf("Query 'CALL EDITProduct(?, ?, ?, ?, ?)', arguments do not match: expected 5, but got 4 arguments"))

	err = product.EditProduct(&p)
	assert.Error(t, err)
	assert.EqualError(t, err, "Query 'CALL EDITProduct(?, ?, ?, ?, ?)', arguments do not match: expected 5, but got 4 arguments")
}

func TestEditProduct_PrepareExec(t *testing.T) {
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
				mock.ExpectPrepare("CALL EditProduct").
					WillReturnError(err)
			},
		},
		{
			err: fmt.Errorf("execution error"),
			mockFn: func(err error) {
				mock.ExpectPrepare("CALL EditProduct").ExpectQuery().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	p := viewmodels.Product{
		ID:          1,
		Name:        "Burger",
		Description: "Burger",
		Price:       10.0,
		Quantity:    10,
	}

	for _, data := range testData {
		data.mockFn(data.err)
		err = product.EditProduct(&p)
		assert.NotNil(t, err, "expected error to not be nil, got %v", err)
		assert.Equal(t, data.err, err, "expected error to be %v, got %v", data.err, err)
		assert.Nil(t, mock.ExpectationsWereMet(), "expected all expectations to be met, got %v", mock.ExpectationsWereMet())
	}
}
