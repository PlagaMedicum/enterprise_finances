package postgresql

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // needed for postgresql db migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // needed for postgresql db migrations.
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // needed for postgresql connection.
	log "github.com/sirupsen/logrus"
)

const (
	migrationPath = "file://server/db/postgresql/migrations"
)

type DB struct {
	User         string
	Password     string
	Host         string
	Port         uint64
	DatabaseName string
	SSLMode      string
	DB           *sqlx.DB
	m            *migrate.Migrate
}

// Connect ...
func (db *DB) Connect() {
	log.Info("Connecting to PostgreSQL DB...")

	var err error
	db.DB, err = sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=%s",
			db.User, db.Password, db.Host, db.Port, db.DatabaseName, db.SSLMode))
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("PostgreSQL DB connected!")
}

func (db *DB) createMigration() {
	driver, err := postgres.WithInstance(db.DB.DB, &postgres.Config{DatabaseName: db.DatabaseName})
	if err != nil {
		log.Error(err)
		return
	}

	db.m, err = migrate.NewWithDatabaseInstance(migrationPath, db.DatabaseName, driver)
	if err != nil {
		log.Error(err)
		return
	}
}

// MigrateUp ...
func (db *DB) MigrateUp() {
	if db.m == nil {
		db.createMigration()
	}

	err := db.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Error(err)
	}
}

// MigrateDown ...
func (db *DB) MigrateDown() {
	if db.m == nil {
		db.createMigration()
	}

	err := db.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Error(err)
	}
}
