package usecases

import (
	"context"
	grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
	"math/big"
)

type Controller struct {
	Repository
}

func AddInfo(ctx context.Context, g grade.Grade) (big.Int, error) {
	return big.Int{}, nil
}

func EditInfo(ctx context.Context) error {
	return nil
}

func DeleteInfo(ctx context.Context) error {
	return nil
}

func GetGradeList(ctx context.Context) error {
	return nil
}
