package grpc

import (
    "context"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/api"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/handlers"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
    "sync"
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

func messageToGrade(g *api.Grade) (domain.Grade, error) {
    d, err := unmarshalDate(g.GetDate())
    if err != nil {
        return domain.Grade{}, err
    }

    return domain.Grade{ID: g.GetId(), Num: g.GetNum(), Date: d, Coefficient: int(g.GetCoeff())}, nil
}

func gradeToMessage(g domain.Grade) *api.Grade {
    return &api.Grade{Id: g.ID, Num: g.Num, Date: g.Date.String(), Coeff: int32(g.Coefficient)}
}

func sendGrades(glist []domain.Grade, send func(*api.Grade) error) error {
    errc := make(chan error)
    waitc := make(chan struct{})
    wg := sync.WaitGroup{}
    for _, g := range glist {
        wg.Add(1)
        go func(g domain.Grade) {
            err := send(gradeToMessage(g))
            if err != nil {
                err = errors.Wrap(err, "Error sending grade")
                log.Error(err)
                errc <- err
                return
            }
            wg.Done()
        }(g)
    }

    go func() {
        wg.Wait()
        close(waitc)
    }()

    select {
    case err := <-errc:
        return err
    case <-waitc:
    }

    return nil
}

func (c Controller) AddInfo(ctx context.Context, rq *api.Grade) (*empty.Empty, error) {
    g, err := messageToGrade(rq)
    if err != nil {
        return &empty.Empty{}, err
    }

    err = c.Usecases.AddInfo(ctx, g)
    if err != nil {
        err = errors.Wrap(err, "Error adding info")
        log.Error(err)
        return &empty.Empty{}, err
    }

    return &empty.Empty{}, nil
}

func (c Controller) EditInfo(ctx context.Context, rq *api.Grade) (*empty.Empty, error) {
    g, err := messageToGrade(rq)
    if err != nil {
        return &empty.Empty{}, err
    }

    err = c.Usecases.EditInfo(ctx, g)
    if err != nil {
        err = errors.Wrap(err, "Error editing info")
        log.Error(err)
        return &empty.Empty{}, err
    }

    return &empty.Empty{}, nil
}

func (c Controller) DeleteInfo(ctx context.Context, id *api.ID) (*empty.Empty, error) {
    err := c.Usecases.DeleteInfo(ctx, id.GetId())
    if err != nil {
        err = errors.Wrap(err, "Error deleting info")
        log.Error(err)
        return &empty.Empty{}, err
    }

    return &empty.Empty{}, nil
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

    err = sendGrades(glist, resp.Send)
    if err != nil {
        return err
    }

    return nil
}

func (c Controller) GetGrade(ctx context.Context, id *api.ID) (*api.Grade, error) {
    g, err := c.Usecases.GetGrade(ctx, id.GetId())
    if err != nil {
        err = errors.Wrap(err, "Error getting grade list")
        log.Error(err)
        return &api.Grade{}, err
    }

    return gradeToMessage(g), nil
}
