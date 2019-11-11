package handlers

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"math/big"
)

type Usecases interface {
	AddInfo(ctx context.Context, g grade.Grade) (big.Int, error)
	EditInfo(ctx context.Context, g grade.Grade) error
	DeleteInfo(ctx context.Context, id big.Int) error
	GetGradeList(ctx context.Context) ([]grade.Grade, error)
}
