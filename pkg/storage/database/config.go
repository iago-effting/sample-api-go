package database

import (
	"database/sql"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type service struct {
	DatabaseOptions DatabaseOptions
	Logger          log.Logger
}

type Service interface {
	Connect() error
	GetDb() *sql.DB
}

type DatabaseOptions struct {
	DSN string
}

var db *bun.DB

func NewDatabaseService(options DatabaseOptions, logger log.Logger) Service {
	return &service{
		DatabaseOptions: options,
		Logger:          logger,
	}
}

func (s service) GetDb() *sql.DB {
	return db.DB
}

func (s service) Connect() error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(s.DatabaseOptions.DSN)))

	db = bun.NewDB(sqldb, pgdialect.New())
	ping := db.Ping()

	if ping != nil {
		level.Debug(s.Logger).Log("DatabaseConnect", ping.Error())
	}

	return nil
}
