package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"

	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/storage/database"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowError())
	}

	configService := configs.NewConfigService(os.Getenv("ENV"), logger)
	configService.LoadEnvVars()

	_, err := database.StartConnection()
	if err != nil {
		level.Error(logger).Log(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", configs.Env.Migrations.Dir),
		configs.Env.Database.DSN,
	)

	if err != nil {
		level.Error(logger).Log(err)
	}

	app := &cli.App{
		Name:  "Focus",
		Usage: "Manager your app",
		Commands: []*cli.Command{
			MigrationCommands(m, logger),
			MakeCommands(m, logger),
		},
	}

	if err := app.Run(os.Args); err != nil {
		level.Error(logger).Log(err)
	}
}

func MakeCommands(migrator *migrate.Migrate, logger log.Logger) *cli.Command {
	return &cli.Command{
		Name:  "make",
		Usage: "Manager your make actions",
		Subcommands: []*cli.Command{
			{
				Name:  "migration",
				Usage: "Create a new migration",
				Action: func(ctx *cli.Context) error {
					var stderr bytes.Buffer
					var out bytes.Buffer

					args := ctx.Args()

					cmd := exec.Command("./bin/migrate", "create", "-dir", configs.Env.Migrations.Dir, "-ext", "sql", args.First())
					fmt.Println(cmd.String())

					cmd.Stderr = &stderr
					cmd.Stdout = &out

					if err := cmd.Run(); err != nil {
						level.Error(logger).Log(err)
						return nil
					}

					return nil
				},
			},
		},
	}
}

func MigrationCommands(migrator *migrate.Migrate, logger log.Logger) *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Manager your database migrations",
		Action: func(ctx *cli.Context) error {
			if err := migrator.Up(); err != nil {
				level.Error(logger).Log("error", err.Error())
				return err
			}

			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "reset",
				Usage: "Rollback all migrations",
				Action: func(c *cli.Context) error {
					if err := migrator.Down(); err != nil {
						level.Error(logger).Log(err)
						return err
					}

					return nil
				},
			},
		},
	}
}
