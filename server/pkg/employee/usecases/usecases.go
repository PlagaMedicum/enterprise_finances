package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
)

type Controller struct {
	Repository
}

// AddEmployee ...
func (c Controller) AddEmployee(ctx context.Context, e employee.Employee) (uint64, error) {
	id, err := c.Repository.AddEmployee(ctx, e)
	return id, err
}

// EditEmployee ...
func (c Controller) EditEmployee(ctx context.Context, e employee.Employee) error {
	err := c.Repository.UpdateEmployee(ctx, e)
	return err
}

// DeleteEmployee ...
func (c Controller) DeleteEmployee(ctx context.Context, id uint64) error {
	err := c.Repository.DeleteEmployee(ctx, id)
	return err
}

// GetEmployeeList ...
func (c Controller) GetEmployeeList(ctx context.Context) ([]employee.Employee, error) {
	elist, err := c.Repository.GetEmployeeList(ctx)

	for _, e := range elist {
		go func() {
			e.Salary["sas"] = 1
		}()
	}

	return elist, err
}

// GetEmployee ...
func (c Controller) GetEmployee(ctx context.Context, id uint64) (employee.Employee, error) {
	e, err := c.Repository.GetEmployee(ctx, id)
	return e, err
}
