package main

import (
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/database/mongodb"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/api"
    handler "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/handlers/grpc"
    repository "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/repositories/mongodb"
    "github.com/PlagaMedicum/enterprise_finances/server/pkg/employee/usecases"
    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc"
    "net"
)

func main() {
    db := mongodb.DB{
        Host:         "localhost",
        Port:         "27017",
        DatabaseName: "efinances",
    }
    err := db.Connect()
    if err != nil {
        log.Fatal(err)
    }

    s := grpc.NewServer()
    api.RegisterEmployeesServer(s, &handler.Controller{usecases.Controller{repository.Controller{db}}})

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
