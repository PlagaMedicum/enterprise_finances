package postgresql

import (
	"database/sql"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type DB struct {
	User     string
	Password string
	Host     string
	Port     uint64
	Database string
	sqldb    *sql.DB
	conn     *pgx.Conn
}

func (db DB) Connect() {
	log.Info("Connectiong to PostgreSQL DB.")

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
}
