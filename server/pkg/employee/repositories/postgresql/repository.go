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
	tx, err := c.DB.BeginTx(ctx)
	if err != nil {
		return big.Int{}, err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRowContext(ctx,
		`insert into employee (name, position, tu_membership) values ($1, $2, $3) returning id;`,
		e.Name, e.Position, e.TUMembership).Scan(&e.ID)
	if err != nil {
		return big.Int{}, err
	}
	_, err = tx.ExecContext(ctx,
		`insert into employees_grades (e_id, g_id) values ($1, $2)`,
		e.ID, e.Grade)
	if err != nil {
		return big.Int{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return big.Int{}, err
	}

	return e.ID, nil
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
