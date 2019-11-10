package postgresql

import (
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
)

type Controller struct {
	postgresql.DB
}

func (c Controller) AddInfo() error {
	return nil
}

func (c Controller) UpdateInfo() error {
	return nil
}

func (c Controller) DeleteInfo() error {
	return nil
}

func (c Controller) GetGradeList() error {
	return nil
}
