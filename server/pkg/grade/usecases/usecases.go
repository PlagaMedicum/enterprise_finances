package usecases

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
)

type Controller struct {
	Repository
}

// AddInfo ...
func (c Controller) AddInfo(ctx context.Context, g grade.Grade) (uint64, error) {
	id, err := c.Repository.AddInfo(ctx, g)
	return id, err
}

// EditInfo ...
func (c Controller) EditInfo(ctx context.Context, g grade.Grade) error {
	err := c.Repository.UpdateInfo(ctx, g)
	return err
}

// DeleteInfo ...
func (c Controller) DeleteInfo(ctx context.Context, id uint64) error {
	err := c.Repository.DeleteInfo(ctx, id)
	return err
}

// GetGradeList
func (c Controller) GetGradeList(ctx context.Context) ([]grade.Grade, error) {
	glist, err := c.Repository.GetGradeList(ctx)
	return glist, err
}
