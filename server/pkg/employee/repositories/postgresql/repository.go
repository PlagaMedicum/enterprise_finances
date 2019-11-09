package postgresql

import (
	psql "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
)

type Controller struct {
	psql.DB
}

func (c Controller) AddEmployee() error {

}

func (c Controller) UpdateEmployee() error {

}

func (c Controller) DeleteEmployee() error {

}

func (c Controller) GetEmployeeList() error {

}
