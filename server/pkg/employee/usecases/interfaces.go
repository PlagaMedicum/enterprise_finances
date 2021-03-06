package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"time"
)

type Repository interface {
	AddEmployee(ctx context.Context, e employee.Employee) (uint64, error)
	UpdateEmployee(ctx context.Context, e employee.Employee) error
	DeleteEmployee(ctx context.Context, id uint64) error
	GetEmployeeList(ctx context.Context, d time.Time) ([]employee.Employee, error)
	GetEmployee(ctx context.Context, id uint64) (employee.Employee, error)
	GetEmployeePayments(ctx context.Context, id uint64, d time.Time) ([]employee.Employee, error)
}
