package postgresql

import (
	"context"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"math/big"
)

type Controller struct {
	postgresql.DB
}

func (c Controller) AddEmployee(ctx context.Context, e employee.Employee) (big.Int, error) {
	return big.Int{}, nil
}

func (c Controller) UpdateEmployee(ctx context.Context, e employee.Employee) error {
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
