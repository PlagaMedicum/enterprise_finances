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
	tx, err := c.DB.DB.BeginTx(ctx, nil)
	if err != nil {
		return big.Int{}, err
	}

	err = tx.QueryRowContext(ctx,
		`insert into employees (name, position, tu_membership) values (:name, :position, :tu_membership) returning id;`,
		e).Scan(&e.ID)
	if err != nil {
		return big.Int{}, err
	}
	_, err = tx.ExecContext(ctx,
		`insert into employees_grades (e_id, g_id) values (:id, :grade)`,
		e)
	if err != nil {
		return big.Int{}, err
	}

	err = tx.Commit()
	if err != nil {
		return big.Int{}, err
	}

	return e.ID, nil
}

func (c Controller) UpdateEmployee(ctx context.Context, e employee.Employee) error {
	tx, err := c.DB.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`update employees set name = :name, position = :position, tu_membership = :tu_membership where id = :id;`,
		e)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx,
		`update employees_grades set g_id = :grade where e_id = :id`,
		e)
	if err != nil {
		return err
	}

	return nil
}

func (c Controller) DeleteEmployee(ctx context.Context, id big.Int) error {
	tx, err := c.DB.DB.BeginTx(ctx, nil)
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

func (c Controller) GetEmployeeList(ctx context.Context) ([]employee.Employee, error) {
	return nil, nil
}

func (c Controller) GetEmployee(ctx context.Context, id big.Int) (employee.Employee, error) {
	return employee.Employee{}, nil
}
