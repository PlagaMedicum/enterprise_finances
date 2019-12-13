package postgresql

import (
	"context"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
	"github.com/pkg/errors"
	"time"
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
		`update employees set name = :name, position = :position, tu_membership = :tu_membership where id = :id;
				update employees_grades set g_id = :grade where e_id = :id`,
		e)
	if err != nil {
		return err
	}

	err = tx.Commit()
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
		`delete from employees where id = $1;
				delete from employees_grades where e_id = $1`,
		id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetEmployeeList ...
func (c Controller) GetEmployeeList(ctx context.Context, d time.Time) ([]employee.Employee, error) {
	rows, err := c.DB.DB.QueryContext(ctx,
		`with with_dates as (
    			select e.id, e.name, e.position, e.tu_membership, ms.date, (ms.value * g.coeff) o from minimal_salaries ms, employees e
    			join employees_grades eg on e.id = eg.e_id
    			join grades g on eg.g_id = g.id and g.date <= $1
    			where ms.date >= $1
			)
			select w.id, w.name, w.position, w.date, cast((w.o + 0.15 * w.o - (0.13 + 0.1 + (case when w.tu_membership then 0.1 else 0 end)) * w.o) as float) salary from with_dates w`,
		d)
	if err != nil {
		return nil, errors.Errorf("Error selecting employees rows: %s", err)
	}

	var eList []employee.Employee
	for rows.Next() {
		e := employee.Employee{}

		err = rows.Scan(&e.ID, &e.Name, &e.Position, &e.Date, &e.Salary)
		if err != nil {
			return nil, err
		}

		eList = append(eList, e)
	}

	return eList, nil
}

// GetEmployee ...
func (c Controller) GetEmployee(ctx context.Context, id uint64) (employee.Employee, error) {
	e := employee.Employee{ID: id}

	err := c.DB.DB.QueryRowContext(ctx,
		`select e.name, e.position, e.tu_membership, g.num from employees e
				join employees_grades eg on e.id = eg.e_id
				join grades g on eg.g_id = g.id
				where e.id = $1`,
		id).Scan(&e.Name, &e.Position, &e.TUMembership, &e.Grade)
	if err != nil {
		return employee.Employee{}, err
	}

	return e, nil
}

// GetEmployeePayments ...
func (c Controller) GetEmployeePayments(ctx context.Context, id uint64, d time.Time) ([]employee.Employee, error) {
	rows, err := c.DB.DB.QueryContext(ctx,
		`with with_dates as (
    			select e.id, e.tu_membership, ms.date, (ms.value * g.coeff) o from minimal_salaries ms, employees e
    			join employees_grades eg on e.id = eg.e_id
    			join grades g on eg.g_id = g.id and g.date <= $2
    			where e.id = $1 and ms.date >= $2
			), payments as (
			    select w.date, cast((w.o + 0.15 * w.o) as float) acc, cast(((0.13 + 0.1 + (case when w.tu_membership then 0.1 else 0 end)) * w.o) as float) ded from with_dates w
			)
			select p.date, p.acc, p.ded, cast((p.acc - p.ded) as float) salary from payments p`,
		id, d)
	if err != nil {
		return nil, errors.Errorf("Error selecting payments rows: %s", err)
	}

	var eList []employee.Employee
	for rows.Next() {
		e := employee.Employee{}

		err = rows.Scan(&e.Date, &e.Accruals, &e.Deduction, &e.Salary)
		if err != nil {
			return nil, err
		}

		eList = append(eList, e)
	}

	return eList, nil
}
