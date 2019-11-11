package postgresql

import (
	"context"
	"github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"math/big"
)

type Controller struct {
	postgresql.DB
}

func (c Controller) AddInfo(ctx context.Context, g grade.Grade) (big.Int, error) {
	return big.Int{}, nil
}

func (c Controller) UpdateInfo(ctx context.Context, g grade.Grade) error {
	return nil
}

func (c Controller) DeleteInfo(ctx context.Context, id big.Int) error {
	return nil
}

func (c Controller) GetGradeList(ctx context.Context) ([]grade.Grade, error) {
	return nil, nil
}
