package mongodb

import (
    "context"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/mongodb"
    employee "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/domain"
    "github.com/pkg/errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Controller struct {
    mongodb.DB
}

func (c Controller) AddEmployee(ctx context.Context, e employee.Employee) (uint64, error) {
    _, err := c.EmployeeCol.InsertOne(ctx, e)
    return e.ID, errors.Wrap(err, "Error inserting employee")
}

func (c Controller) UpdateEmployee(ctx context.Context, e employee.Employee) error {
    filter := bson.D{
        {
            Key:   "id",
            Value: e.ID,
        },
    }
    _, err := c.EmployeeCol.UpdateOne(ctx, filter, e)
    return errors.Wrap(err, "Error updating employee")
}

func (c Controller) DeleteEmployee(ctx context.Context, id uint64) error {
    filter := bson.D{
        {
            Key:   "id",
            Value: id,
        },
    }
    _, err := c.EmployeeCol.DeleteOne(ctx, filter)
    return errors.Wrap(err, "Error deleting employee")
}

func (c Controller) GetEmployeeList(ctx context.Context, d time.Time) ([]employee.Employee, error) {
    filter := bson.D{
        {
            Key:   "date",
            Value: d,
        },
    }
    var glist []employee.Employee
    cur, err := c.EmployeeCol.Find(ctx, filter)
    if err != nil {
        return nil, errors.Wrapf(err, "Error finding grades for %s", d.String())
    }

    for cur.Next(ctx) {
        var res bson.M
        err := cur.Decode(&res)
        if err != nil {
            return nil, errors.Wrap(err, "Error decoding result from cursor")
        }

        pd, ok := res["date"].(primitive.DateTime)
        if !ok {
            return nil, errors.Errorf("Error decoding result from cursor: returned \"date\" field's type is %T, not primitive.DateTime", res["date"])
        }
        d := pd.Time()

        glist = append(glist, employee.Employee{
            ID:   uint64(res["id"].(int64)),
            Date: d,
            Name: res["name"].(string),
            Position: res["position"].(string),
            Grade: uint64(res["grade"].(int64)),
            TUMembership: res["tu_membership"].(bool),
            Salary: res["salary"].(float64),
            Accruals: res["accruals"].(float64),
            Deduction: res["deduction"].(float64),
        })
    }

    return glist, nil
}

func (c Controller) GetEmployee(ctx context.Context, id uint64) (employee.Employee, error) {
    filter := bson.D{
        {
            Key:   "id",
            Value: id,
        },
    }
    var e employee.Employee
    err := c.EmployeeCol.FindOne(ctx, filter).Decode(&e)
    return e, errors.Wrap(err, "Error getting employee")
}

func (c Controller) GetEmployeePayments(ctx context.Context, id uint64, d time.Time) ([]employee.Employee, error) {

    return nil, nil
}
