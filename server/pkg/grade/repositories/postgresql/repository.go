package postgresql

import (
	"context"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
)

type Controller struct {
	postgresql.DB
}

// AddInfo ...
func (c Controller) AddInfo(ctx context.Context, g grade.Grade) (uint64, error) {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	err = tx.QueryRowContext(ctx,
		`insert into grades (date, coeff) values ($1, $2) returning id`,
		g.Date, g.Coefficient).Scan(&g.ID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// UpdateInfo ...
func (c Controller) UpdateInfo(ctx context.Context, g grade.Grade) error {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx,
		`update grades set date = :date, coeff = :coeff where id = :id`,
		g)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// DeleteInfo ...
func (c Controller) DeleteInfo(ctx context.Context, id uint64) error {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`delete from grades where id = $1`,
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

// GetGradeList ...
func (c Controller) GetGradeList(ctx context.Context) ([]grade.Grade, error) {
	rows, err := c.DB.DB.QueryContext(ctx,
		`select (id, date, coeff) from grades`)
	if err != nil {
		return nil, err
	}

	var gList []grade.Grade
	for rows.Next() {
		g := grade.Grade{}
		err = rows.Scan(&g.ID, &g.Date, &g.Coefficient)
		if err != nil {
			return nil, err
		}
		gList = append(gList, g)
	}

	return gList, nil
}
