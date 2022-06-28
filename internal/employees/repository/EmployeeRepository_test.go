package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/repository"
)

func TestEmployeeRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()

	expectedEmployees := []domain.Employee{
		{
			Id:           1,
			CardNumberId: "123456",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		},
		{
			Id:           2,
			CardNumberId: "789012",
			FirstName:    "Jane",
			LastName:     "Doe",
			WarehouseId:  2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
	for _, employee := range expectedEmployees {
		rows = rows.AddRow(employee.Id, employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)
	}

	employeeRepository := repository.NewMariaDBEmployeeRepository(db)

	t.Run("should return all employees", func(t *testing.T) {
		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnRows(rows)

		result, err := employeeRepository.GetAll(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expectedEmployees, result)
	})

	t.Run("should return error when query fails", func(t *testing.T) {
		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnError(fmt.Errorf("query error"))

		result, err := employeeRepository.GetAll(ctx)
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("should return error when scan fails", func(t *testing.T) {
		rows = rows.AddRow("invalid_id_format", "123456", "John", "Doe", 1)
		mock.ExpectQuery(repository.SQLFindAllEmployees).WillReturnRows(rows)

		result, err := employeeRepository.GetAll(ctx)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
