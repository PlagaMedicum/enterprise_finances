package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"math/big"
)

type Repository interface {
	AddEmployee(ctx context.Context, e employee.Employee) (big.Int, error)
	UpdateEmployee(ctx context.Context, e employee.Employee) error
	DeleteEmployee(ctx context.Context, id big.Int) error
	GetEmployeeList(ctx context.Context) ([]employee.Employee, error)
	GetEmployee(ctx context.Context, id big.Int) (employee.Employee, error)
}
