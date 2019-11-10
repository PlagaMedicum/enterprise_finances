package postgresql

import (
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
)

type Controller struct {
	postgresql.DB
}

func (c Controller) AddEmployee() error {
	return nil
}

func (c Controller) UpdateEmployee() error {
	return nil
}

func (c Controller) DeleteEmployee() error {
	return nil
}

func (c Controller) GetEmployeeList() error {
	return nil
}
