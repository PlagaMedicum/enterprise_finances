package grpc

import (
    "context"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/api"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/handlers"
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

func (c Controller) AddEmployee(ctx context.Context, e *api.Employee) (*api.ID, error) {
    d, err := unmarshalDate(e.GetDate())
    if err != nil {
        return nil, err
    }

    id, err := c.Usecases.AddEmployee(ctx, domain.Employee{ID: e.GetId(), Date: d, Name: e.GetName(), Position: e.GetPosition(), Grade: e.GetGrade(), TUMembership: e.GetTuMembership()})
    if err != nil {
        err = errors.Wrap(err, "Error adding employee")
        log.Error(err)
        return nil, err
    }

    return &api.ID{Id: id}, nil
}

func (c Controller) EditEmployee(ctx context.Context, e *api.Employee) (*empty.Empty, error) {
    d, err := unmarshalDate(e.GetDate())
    if err != nil {
        return nil, err
    }

    err = c.Usecases.EditEmployee(ctx, domain.Employee{ID: e.GetId(), Date: d, Name: e.GetName(), Position: e.GetPosition(), Grade: e.GetGrade(), TUMembership: e.GetTuMembership()})
    if err != nil {
        err = errors.Wrap(err, "Error editing employee")
        log.Error(err)
        return nil, err
    }

    return nil, nil
}

func (c Controller) DeleteEmployee(ctx context.Context, id *api.ID) (*empty.Empty, error) {
    err := c.Usecases.DeleteEmployee(ctx, id.GetId())
    if err != nil {
        err = errors.Wrap(err, "Error deleting employee")
        log.Error(err)
        return nil, err
    }

    return nil, nil
}

func (c Controller) GetEmployeeList(date *api.Date, resp api.Employees_GetEmployeeListServer) error {
    d, err := unmarshalDate(date.GetDate())
    if err != nil {
        return err
    }

    elist, err := c.Usecases.GetEmployeeList(resp.Context(), d)
    if err != nil {
        err = errors.Wrap(err, "Error getting employee list")
        log.Error(err)
        return err
    }
    for _, e := range elist {
        err = resp.Send(&api.Employee{Id: e.ID, Date: e.Date.String(), Name: e.Name, Position: e.Position, Grade: e.Grade, TuMembership: e.TUMembership})
        if err != nil {
            err = errors.Wrap(err, "Error sending employee")
            log.Error(err)
            return err
        }
    }

    return nil
}

func (c Controller) GetEmployee(ctx context.Context, id *api.ID) (*api.Employee, error) {
    e, err := c.Usecases.GetEmployee(ctx, id.GetId())
    if err != nil {
        err = errors.Wrap(err, "Error getting employee")
        log.Error(err)
        return nil, err
    }

    return &api.Employee{Id: e.ID, Date: e.Date.String(), Name: e.Name, Position: e.Position, Grade: e.Grade, TuMembership: e.TUMembership}, nil
}

func (c Controller) GetEmployeePayments(rq *api.IDTime, resp api.Employees_GetEmployeePaymentsServer) error {
    d, err := unmarshalDate(rq.GetDate())
    if err != nil {
        return err
    }

    elist, err := c.Usecases.GetEmployeePayments(resp.Context(), rq.GetId(), d)
    if err != nil {
        err = errors.Wrap(err, "Error getting employee list")
        log.Error(err)
        return err
    }
    for _, e := range elist {
        err = resp.Send(&api.Employee{Id: e.ID, Date: e.Date.String(), Name: e.Name, Position: e.Position, Grade: e.Grade, TuMembership: e.TUMembership})
        if err != nil {
            err = errors.Wrap(err, "Error sending employee")
            log.Error(err)
            return err
        }
    }

    return nil
}
