package postgresql

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // needed for postgresql db migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // needed for postgresql db migrations.
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // needed for postgresql connection.
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	migrationPath = "file://db/postgresql/migrations"
)

type DB struct {
	User     string
	Password string
	Host     string
	Port     uint64
	Database string
	SSLMode  string
	DB       *sqlx.DB
	m        *migrate.Migrate
}

func (db *DB) Connect() {
	log.Info("Connecting to PostgreSQL DB.")

	sqlxdb, err := sqlx.Connect("postgres",
		"user="+db.User+" "+
			"password="+db.Password+" "+
			"host="+db.Host+" "+
			"port="+strconv.FormatUint(db.Port, 10)+" "+
			"database="+db.Database+" "+
			"sslmode="+db.SSLMode)
	if err != nil {
		log.Error(err)
		return
	}

	err = sqlxdb.Ping()
	if err != nil {
		log.Error(err)
		return
	}

	db.DB = sqlxdb
	sqlxdb.BeginTx()

	log.Info("PostgreSQL DB connected!")
}

func (db *DB) createMigration() {
	driver, err := postgres.WithInstance(db.DB.DB, &postgres.Config{DatabaseName: db.Database})
	if err != nil {
		log.Error(err)
		return
	}

	db.m, err = migrate.NewWithDatabaseInstance(migrationPath, db.Database, driver)
	if err != nil {
		log.Error(err)
		return
	}
}

func (db *DB) MigrateUp() {
	if db.m == nil {
		db.createMigration()
	}

	err := db.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Error(err)
	}
}

func (db *DB) MigrateDown() {
	if db.m == nil {
		db.createMigration()
	}

	err := db.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Error(err)
	}
}
