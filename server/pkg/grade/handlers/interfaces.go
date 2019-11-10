package handlers

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"math/big"
)

type Usecases interface {
	AddInfo(ctx context.Context, g grade.Grade) (big.Int, error)
	EditInfo(ctx context.Context) error
	DeleteInfo(ctx context.Context) error
	GetGradeList(ctx context.Context) error
}
