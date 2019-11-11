package usecases

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"math/big"
)

type Controller struct {
	Repository
}

func (c Controller) AddInfo(ctx context.Context, g grade.Grade) (big.Int, error) {
	id, err := c.Repository.AddInfo(ctx, g)
	return id, err
}

func (c Controller) EditInfo(ctx context.Context, g grade.Grade) error {
	err := c.Repository.UpdateInfo(ctx, g)
	return err
}

func (c Controller) DeleteInfo(ctx context.Context, id big.Int) error {
	err := c.Repository.DeleteInfo(ctx, id)
	return err
}

func (c Controller) GetGradeList(ctx context.Context) ([]grade.Grade, error) {
	glist, err := c.Repository.GetGradeList(ctx)
	return glist, err
}
