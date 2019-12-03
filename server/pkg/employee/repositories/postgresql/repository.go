package postgresql

import (
	"context"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"github.com/pkg/errors"
)

type Controller struct {
	postgresql.DB
}

// AddEmployee ...
func (c Controller) AddEmployee(ctx context.Context, e employee.Employee) (uint64, error) {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	err = tx.QueryRowxContext(ctx,
		`insert into employees (name, position, tu_membership) values ($1, $2, $3) returning id;`,
		e.Name, e.Position, e.TUMembership).Scan(&e.ID)
	if err != nil {
		return 0, errors.Errorf("Query row error. \"insert into employees...\": %s", err)
	}
	_, err = tx.NamedExecContext(ctx,
		`insert into employees_grades (e_id, g_id) values (:id, :grade)`,
		e)
	if err != nil {
		return 0, errors.Errorf("Exec error. \"insert into employees_grades...\": %s", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return e.ID, nil
}

// UpdateEmployee ...
func (c Controller) UpdateEmployee(ctx context.Context, e employee.Employee) error {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx,
		`update employees set name = :name, position = :position, tu_membership = :tu_membership where id = :id;`,
		e)
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(ctx,
		`update employees_grades set g_id = :grade where e_id = :id`,
		e)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEmployee ...
func (c Controller) DeleteEmployee(ctx context.Context, id uint64) error {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`delete from employees where id = $1`,
		id)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx,
		`delete from employees_grades where e_id = $1`,
		id)

	return nil
}

func (c Controller) GetEmployeeCoefficient(ctx context.Context) (uint64, error) { // TODO: Specify the date
	// TODO
	return 0, nil
}

// GetEmployeeList ...
func (c Controller) GetEmployeeList(ctx context.Context) ([]employee.Employee, error) { // TODO: Specify the date
	rows, err := c.DB.DB.QueryContext(ctx,
		`select (id, name, position, tu_membership) from employees`)
	if err != nil {
		return nil, errors.Errorf("Error selecting employees rows: %s", err)
	}

	var eList []employee.Employee
	for rows.Next() {
		e := employee.Employee{}
		err = rows.Scan(&e.ID, &e.Name, &e.Position, &e.TUMembership) // FIXME: Fix multiple scan
		if err != nil {
			return nil, err
		}
		err = c.DB.DB.QueryRowContext(ctx,
			`select g_id from employees_grades where e_id = $1`,
			e.ID).Scan(&e.Grade)
		if err != nil {
			return nil, err
		}
		// TODO: Calculate salary for the date
		eList = append(eList, e)
	}

	return eList, nil
}

// GetEmployee ...
func (c Controller) GetEmployee(ctx context.Context, id uint64) (employee.Employee, error) { // TODO: Specify the date
	e := employee.Employee{ID: id}

	err := c.DB.DB.QueryRowContext(ctx,
		`select (name, position, tu_membership) from employees where id = $1`,
		id).Scan(&e.Name, &e.Position, &e.TUMembership) // FIXME: Fix multiple scan
	if err != nil {
		return employee.Employee{}, err
	}
	err = c.DB.DB.QueryRowContext(ctx,
		`select g_id from employees_grades where e_id = $1`,
		id).Scan(&e.Grade)
	if err != nil {
		return employee.Employee{}, err
	}
	// TODO: Calculate salary for the date
	return e, nil
}
