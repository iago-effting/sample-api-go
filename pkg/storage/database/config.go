package database

import (
	"database/sql"
	"fmt"

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
	GetDb() *bun.DB
}

type DatabaseOptions struct {
	DSN      string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

var BunDb *bun.DB

func NewDatabaseService(options DatabaseOptions, logger log.Logger) Service {
	return &service{
		DatabaseOptions: options,
		Logger:          logger,
	}
}

func (s service) GetDb() *bun.DB {
	return BunDb
}

func (s service) Connect() error {
	if s.DatabaseOptions.DSN == "" {
		s.DatabaseOptions.DSN = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			s.DatabaseOptions.User,
			s.DatabaseOptions.Password,
			s.DatabaseOptions.Host,
			s.DatabaseOptions.Port,
			s.DatabaseOptions.Name,
		)
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(s.DatabaseOptions.DSN)))

	BunDb = bun.NewDB(sqldb, pgdialect.New())
	ping := BunDb.Ping()

	if ping != nil {
		level.Debug(s.Logger).Log("DatabaseConnect", ping.Error())
	}

	return nil
}
