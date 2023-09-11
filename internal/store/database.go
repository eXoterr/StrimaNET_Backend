package store

import (
	"errors"
	"fmt"
	"os"

	"github.com/eXoterr/StrimaNET_Backend/internal/store/postgres/user"
	"github.com/eXoterr/StrimaNET_Backend/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	Configure() error
	Connect() error

	User() user.UserStore
}

type PgDatabase struct {
	DB         *sqlx.DB
	connString string
}

func (pgdb *PgDatabase) Connect() error {
	db, err := sqlx.Connect("postgres", pgdb.connString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	logger.GetLogger().Debug(logger.LoggingData{
		Data: "connected to db",
	})

	pgdb.DB = db

	return nil
}

func (pgdb *PgDatabase) Configure() error {
	connString := "postgres://%s:%s@%s/%s?sslmode=%s"

	name := os.Getenv("DB_NAME")
	if name == "" {
		return errors.New("db name is empty, check DB_NAME env var")
	}

	password := os.Getenv("DB_PASSWORD")

	host := os.Getenv("DB_HOST")
	if host == "" {
		return errors.New("db host is empty, check DB_HOST env var")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return errors.New("db user is empty, check DB_USER env var")
	}

	ssl := os.Getenv("DB_SSL")
	if ssl == "" {
		return errors.New("db ssl is empty, check DB_SSL env var")
	}

	connString = fmt.Sprintf(
		connString,
		user,
		password,
		host,
		name,
		ssl,
	)

	pgdb.connString = connString

	return nil
}
