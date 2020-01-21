package mongodb

import (
    "context"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
    Host         string
    Port         string
    DatabaseName string
    EmployeeCol  *mongo.Collection
    GradeCol     *mongo.Collection
    Client       *mongo.Client
}

func (db *DB) Connect() error {
    log.Info("Connecting to MongoDB...")

    opt := options.Client().ApplyURI("mongodb://" + db.Host + ":" + db.Port)
    var err error
    db.Client, err = mongo.NewClient(opt)
    if err != nil {
        return errors.Wrap(err, "Failed creating new MongoDB client")
    }

    ctx := context.Background()
    err = db.Client.Connect(ctx)
    if err != nil {
        return errors.Wrap(err, "Failed connecting to MongoDB")
    }

    err = db.Client.Ping(ctx, nil)
    if err != nil {
        return errors.Wrap(err, "Failed to ping MongoDB")
    }

    db.EmployeeCol = db.Client.Database(db.DatabaseName).Collection("employeeCollection")
    db.GradeCol = db.Client.Database(db.DatabaseName).Collection("gradeCollection")

    log.Info("Connected to MongoDB!")
    return nil
}
