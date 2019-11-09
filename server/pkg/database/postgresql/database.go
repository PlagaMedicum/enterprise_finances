package postgresql

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // needed for postgresql db migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // needed for postgresql db migrations.
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq" // needed for postgresql db migrations.
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
	sqldb    *sql.DB
	conn     *pgx.Conn
	m        *migrate.Migrate
}

func (db *DB) Connect() {
	log.Info("Connecting to PostgreSQL DB.")

	var err error
	db.sqldb, err = sql.Open("pgx",
		"user="+db.User+" "+
			"password="+db.Password+" "+
			"host="+db.Host+" "+
			"port="+strconv.FormatUint(db.Port, 10)+" "+
			"database="+db.Database)
	if err != nil {
		log.Error(err)
	}

	err = db.sqldb.Ping()
	if err != nil {
		log.Error(err)
	}

	db.conn, err = stdlib.AcquireConn(db.sqldb)
	if err != nil {
		log.Error(err)
	}

	log.Info("PostgreSQL DB connected!")
}

func (db *DB) createMigration() {
	driver, err := postgres.WithInstance(db.sqldb, &postgres.Config{DatabaseName: db.Database})
	if err != nil {
		log.Error(err)
	}

	db.m, err = migrate.NewWithDatabaseInstance(migrationPath, db.Database, driver)
	if err != nil {
		log.Error(err)
	}
}

func (db *DB) MigrateUp() {
	err := db.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Error(err)
	}
}

func (db *DB) MigrateDown() {
	err := db.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Error(err)
	}
}
