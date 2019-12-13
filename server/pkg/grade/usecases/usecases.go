package usecases

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"github.com/pkg/errors"
	"time"
)

type Controller struct {
	Repository
}

// AddInfo ...
func (c Controller) AddInfo(ctx context.Context, g grade.Grade) error {
	if g.Num > 18 {
		return errors.Errorf("Grade's order number can't be grater than 18.")
	}

	err := c.Repository.AddInfo(ctx, g)
	return err
}

// EditInfo ...
func (c Controller) EditInfo(ctx context.Context, g grade.Grade) error {
	if g.Num > 18 {
		return errors.Errorf("Grade's order number can't be grater than 18.")
	}

	err := c.Repository.UpdateInfo(ctx, g)
	return err
}

// DeleteInfo ...
func (c Controller) DeleteInfo(ctx context.Context, id uint64) error {
	err := c.Repository.DeleteInfo(ctx, id)
	return err
}

// GetGradeList ...
func (c Controller) GetGradeList(ctx context.Context, d time.Time) ([]grade.Grade, error) {
	glist, err := c.Repository.GetGradeList(ctx, d)
	return glist, err
}

// GetGrade ...
func (c Controller) GetGrade(ctx context.Context, id uint64) (grade.Grade, error) {
	g, err := c.Repository.GetGrade(ctx, id)
	return g, err
}
