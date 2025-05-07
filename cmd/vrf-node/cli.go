package main

import (
	"context"
	VRF_Node "github.com/cozyxiong/VRF-Node"

	"github.com/ethereum/go-ethereum/log"

	"github.com/urfave/cli/v2"

	"github.com/cozyxiong/VRF-Node/common/cliapp"
	"github.com/cozyxiong/VRF-Node/common/opio"
	"github.com/cozyxiong/VRF-Node/config"
	"github.com/cozyxiong/VRF-Node/database"
	flag2 "github.com/cozyxiong/VRF-Node/flags"
)

func runVrfNode(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("Running vrf node...")
	cfg := config.NewConfig(ctx)

	return VRF_Node.NewVrfNode(ctx.Context, &cfg, shutdown)
}

func runMigrations(ctx *cli.Context) error {
	log.Info("Running migrations...")
	cfg := config.NewConfig(ctx)

	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("Failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	return db.ExecuteSQLMigration(cfg.Migrations)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flag2.Flags
	return &cli.App{
		Version:              "v0.0.1",
		Description:          "An indexer of events with a serving api",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      cliapp.LifecycleCmd(runVrfNode),
			},
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Runs the database migrations service",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "Print the version of VRF-Node",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
