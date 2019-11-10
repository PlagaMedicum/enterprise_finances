package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"math/big"
)

type Controller struct {
	Repository
}

func (c Controller) AddEmployee(ctx context.Context) (big.Int, error) {
	return big.Int{}, nil
}

func (c Controller) EditEmployee(ctx context.Context, id big.Int) error {
	return nil
}

func (c Controller) DeleteEmployee(ctx context.Context, id big.Int) error {
	return nil
}

func (c Controller) GetEmployeeList(ctx context.Context) ([]employee.Employee, error) {
	return nil, nil
}

func (c Controller) GetEmployee(ctx context.Context, id big.Int) (employee.Employee, error) {
	return employee.Employee{}, nil
}
