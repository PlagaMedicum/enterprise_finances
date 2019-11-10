package handlers

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"math/big"
)

type Usecases interface {
	AddEmployee(ctx context.Context, e employee.Employee) (big.Int, error)
	EditEmployee(ctx context.Context, id big.Int) error
	DeleteEmployee(ctx context.Context, id big.Int) error
	GetEmployeeList(ctx context.Context) ([]employee.Employee, error)
}
