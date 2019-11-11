package usecases

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"math/big"
)

type Repository interface {
	AddInfo(ctx context.Context, g grade.Grade) (big.Int, error)
	UpdateInfo(ctx context.Context, g grade.Grade) error
	DeleteInfo(ctx context.Context, id big.Int) error
	GetGradeList(ctx context.Context) ([]grade.Grade, error)
}
