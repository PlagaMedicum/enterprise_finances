package mongodb

import (
    "context"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/mongodb"
    grade "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/domain"
    "github.com/pkg/errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Controller struct {
    mongodb.DB
}

func (c Controller) AddInfo(ctx context.Context, g grade.Grade) error {
    _, err := c.GradeCol.InsertOne(ctx, g)
    return errors.Wrap(err, "Error inserting grade")
}

func (c Controller) UpdateInfo(ctx context.Context, g grade.Grade) error {
    filter := bson.D{
        {
            Key: "id",
            Value: g.ID,
        },
    }
    _, err := c.GradeCol.UpdateOne(ctx, filter, g)
    return errors.Wrap(err, "Error updating grade")
}

func (c Controller) DeleteInfo(ctx context.Context, id uint64) error {
    filter := bson.D{
        {
            Key: "id",
            Value: id,
        },
    }
    _, err := c.GradeCol.DeleteOne(ctx, filter)
    return errors.Wrap(err, "Error deleting grade")
}

func (c Controller) GetGradeList(ctx context.Context, d time.Time) ([]grade.Grade, error) {
    filter := bson.D{
        {
            Key: "date",
            Value: d,
        },
    }
    var glist []grade.Grade
    cur, err := c.GradeCol.Find(ctx, filter)
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

        glist = append(glist, grade.Grade{
            ID:          uint64(res["id"].(int64)),
            Num:         uint64(res["num"].(int64)),
            Date:        d,
            Coefficient: int(res["coeff"].(int32)),
        })
    }

    return glist, nil
}

func (c Controller) GetGrade(ctx context.Context, id uint64) (grade.Grade, error) {
    filter := bson.D{
        {
            Key: "id",
            Value: id,
        },
    }
    var g grade.Grade
    err := c.GradeCol.FindOne(ctx, filter).Decode(&g)
    return g, errors.Wrap(err, "Error getting grade")
}
