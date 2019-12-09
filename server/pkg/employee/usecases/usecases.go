package usecases

import (
	"context"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"sync"
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

	var w sync.WaitGroup
	for i := range elist { // TODO: Create Salary calculation logic
		elist[i].Salary = make(map[string]uint64)
		w.Add(1)
		go func(e *employee.Employee, w *sync.WaitGroup, i int) {
			e.Salary["sas"] = uint64(i)
			w.Done()
		}(&elist[i], &w, i)
	}

	w.Wait()
	return elist, err
}

// GetEmployee ...
func (c Controller) GetEmployee(ctx context.Context, id uint64) (employee.Employee, error) {
	e, err := c.Repository.GetEmployee(ctx, id)
	return e, err
}
