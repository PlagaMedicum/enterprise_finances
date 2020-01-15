package main

import (
    psql "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/postgresql"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/api"
    handler "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/handlers/grpc"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/repositories/postgresql"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/grade/usecases"
    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc"
    "net"
)

func main() {
    db := psql.DB{
        User:         "postgres",
        Password:     "postgres",
        Host:         "localhost",
        Port:         5432,
        DatabaseName: "efinances",
        SSLMode:      "disable",
    }
    db.Connect()
    db.MigrateDown()
    db.MigrateUp()

    s := grpc.NewServer()
    api.RegisterGradesServer(s, &handler.Controller{usecases.Controller{postgresql.Controller{db}}})

    l, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatal(err)
    }
    log.Info("Listening to localhost:9000")

    err = s.Serve(l)
    if err != nil {
        log.Fatal(err)
    }
}
