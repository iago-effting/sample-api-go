package database

import (
	"database/sql"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
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
	Driver string
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
	sqldb, err := sql.Open(sqliteshim.ShimName, s.DatabaseOptions.Driver)
	if err != nil {
		level.Error(s.Logger).Log("Connect", err)
		return err
	}

	db = bun.NewDB(sqldb, sqlitedialect.New())

	level.Debug(s.Logger).Log("DatabaseConnect", true)

	return nil
}
