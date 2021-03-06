package postgresql

import (
	"context"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"time"
)

type Controller struct {
	postgresql.DB
}

// AddInfo ...
func (c Controller) AddInfo(ctx context.Context, g grade.Grade) error {
	tx, err := c.DB.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx,
		`insert into grades (id, num, date, coeff) values (:id, :num, :date, :coeff) returning id`,
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
func (c Controller) GetGradeList(ctx context.Context, d time.Time) ([]grade.Grade, error) {
	rows, err := c.DB.DB.QueryContext(ctx,
		`select id, num, date, coeff from grades where date >= $1`,
		d)
	if err != nil {
		return nil, err
	}

	var gList []grade.Grade
	for rows.Next() {
		g := grade.Grade{}

		err = rows.Scan(&g.ID, &g.Num, &g.Date, &g.Coefficient)
		if err != nil {
			return nil, err
		}

		gList = append(gList, g)
	}

	return gList, nil
}

// GetGrade ...
func (c Controller) GetGrade(ctx context.Context, id uint64) (grade.Grade, error) {
	g := grade.Grade{ID: id}

	err := c.DB.DB.QueryRowContext(ctx,
		`select num, date, coeff from grades where id = $1`,
		id).Scan(&g.Num, &g.Date, &g.Coefficient)
	if err != nil {
		return grade.Grade{}, err
	}

	return g, nil
}
