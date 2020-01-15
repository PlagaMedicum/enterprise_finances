package grpc

import (
    "context"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/api"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/handlers"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
    "time"
)

type Controller struct {
    handlers.Usecases
}

func unmarshalDate(d string) (time.Time, error) {
    res := time.Time{}
    err := res.UnmarshalText([]byte(d))
    if err != nil {
        err = errors.Wrap(err, "Error unmarshalling date")
        log.Error(err)
        return time.Time{}, err
    }

    return res, nil
}

func (c Controller) AddInfo(ctx context.Context, g *api.Grade) (*empty.Empty, error) {
    d, err := unmarshalDate(g.GetDate())
    if err != nil {
        return nil, err
    }

    err = c.Usecases.AddInfo(ctx, domain.Grade{ID: g.GetId(), Num: g.GetNum(), Date: d, Coefficient: int(g.GetCoeff())})
    if err != nil {
        err = errors.Wrap(err, "Error adding info")
        log.Error(err)
        return nil, err
    }

    return nil, nil
}

func (c Controller) EditInfo(ctx context.Context, g *api.Grade) (*empty.Empty, error) {
    d, err := unmarshalDate(g.GetDate())
    if err != nil {
        return nil, err
    }

    err = c.Usecases.EditInfo(ctx, domain.Grade{ID: g.GetId(), Num: g.GetNum(), Date: d, Coefficient: int(g.GetCoeff())})
    if err != nil {
        err = errors.Wrap(err, "Error editing info")
        log.Error(err)
        return nil, err
    }

    return nil, nil
}

func (c Controller) DeleteInfo(ctx context.Context, id *api.ID) (*empty.Empty, error) {
    err := c.Usecases.DeleteInfo(ctx, id.GetId())
    if err != nil {
        err = errors.Wrap(err, "Error deleting info")
        log.Error(err)
        return nil, err
    }

    return nil, nil
}

func (c Controller) GetGradeList(date *api.Date, resp api.Grades_GetGradeListServer) error {
    d, err := unmarshalDate(date.GetDate())
    if err != nil {
        return err
    }

    glist, err := c.Usecases.GetGradeList(resp.Context(), d)
    if err != nil {
        err = errors.Wrap(err, "Error getting grade list")
        log.Error(err)
        return err
    }
    for _, g := range glist {
        err = resp.Send(&api.Grade{Id: g.ID, Num: g.Num, Date: g.Date.String(), Coeff: int32(g.Coefficient)})
        if err != nil {
            err = errors.Wrap(err, "Error sending grade")
            log.Error(err)
            return err
        }
    }

    return nil
}

func (c Controller) GetGrade(ctx context.Context, id *api.ID) (*api.Grade, error) {
    g, err := c.Usecases.GetGrade(ctx, id.GetId())
    if err != nil {
        err = errors.Wrap(err, "Error getting grade list")
        log.Error(err)
        return nil, err
    }

    return &api.Grade{Id: g.ID, Num: g.Num, Date: g.Date.String(), Coeff: int32(g.Coefficient)}, nil
}
