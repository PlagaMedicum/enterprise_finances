package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
)

type Repository interface {
	AddEmployee(ctx context.Context, e employee.Employee) (uint64, error)
	UpdateEmployee(ctx context.Context, e employee.Employee) error
	DeleteEmployee(ctx context.Context, id uint64) error
	GetEmployeeList(ctx context.Context) ([]employee.Employee, error)
	GetEmployee(ctx context.Context, id uint64) (employee.Employee, error)
}
