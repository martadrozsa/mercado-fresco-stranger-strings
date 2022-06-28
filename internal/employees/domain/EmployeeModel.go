package domain

import "context"

type Employee struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int64  `json:"warehouse_id"`
}

type EmployeeService interface {
	GetAll() ([]Employee, error)
	GetById(id int64) (*Employee, error)
	Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	UpdateFullname(id int64, firstName string, lastName string) (*Employee, error)
	Delete(id int64) error
}

type EmployeeRepository interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int64) (Employee, error)
	UpdateFullname(ctx context.Context, id int64, firstName string, lastName string) (*Employee, error)
	Delete(ctx context.Context, id int64) error
}
