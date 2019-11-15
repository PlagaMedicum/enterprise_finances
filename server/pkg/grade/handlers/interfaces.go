package handlers

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
)

type Usecases interface {
	AddInfo(ctx context.Context, g grade.Grade) (uint64, error)
	EditInfo(ctx context.Context, g grade.Grade) error
	DeleteInfo(ctx context.Context, id uint64) error
	GetGradeList(ctx context.Context) ([]grade.Grade, error)
}
