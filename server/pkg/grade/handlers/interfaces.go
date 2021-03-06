package handlers

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"time"
)

type Usecases interface {
	AddInfo(ctx context.Context, g grade.Grade) error
	EditInfo(ctx context.Context, g grade.Grade) error
	DeleteInfo(ctx context.Context, id uint64) error
	GetGradeList(ctx context.Context, d time.Time) ([]grade.Grade, error)
	GetGrade(ctx context.Context, id uint64) (grade.Grade, error)
}
