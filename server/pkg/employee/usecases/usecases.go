package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"math/big"
)

type Controller struct {
	Repository
}

func (c Controller) AddEmployee(ctx context.Context, e employee.Employee) (big.Int, error) {
	id, err := c.Repository.AddEmployee(ctx, e)
	return id, err
}

func (c Controller) EditEmployee(ctx context.Context, e employee.Employee) error {
	err := c.Repository.UpdateEmployee(ctx, e)
	return err
}

func (c Controller) DeleteEmployee(ctx context.Context, id big.Int) error {
	err := c.Repository.DeleteEmployee(ctx, id)
	return err
}

func (c Controller) GetEmployeeList(ctx context.Context) ([]employee.Employee, error) {
	elist, err := c.Repository.GetEmployeeList(ctx)
	return elist, err
}

func (c Controller) GetEmployee(ctx context.Context, id big.Int) (employee.Employee, error) {
	e, err := c.Repository.GetEmployee(ctx, id)
	return e, err
}
