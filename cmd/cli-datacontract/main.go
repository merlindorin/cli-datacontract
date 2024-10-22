package main

import (
	_ "embed"

	"github.com/alecthomas/kong"
	"github.com/merlindorin/cli-datacontract/cmd/cli-datacontract/commands"
	"github.com/merlindorin/go-shared/pkg/cmd"
)

//nolint:gochecknoglobals // these global variables exist to be overridden during build
var (
	name    = "datacontract"
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

type CLI struct {
	*cmd.Commons
	*cmd.Config

	Bigquery commands.BigqueryCMD `cmd:"bigquery" help:"import datacontract from bigquery" group:"bigquery."`
}

func main() {
	cli := CLI{
		Commons: &cmd.Commons{
			Version: cmd.NewVersion(name, version, commit, buildSource, date),
			Licence: cmd.NewLicence(license),
		},
		Config: cmd.NewConfig(name),
	}

	ctx := kong.Parse(
		&cli,
		kong.Name(name),
		kong.Description("CLI for importing Datacontracts from different sources"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Commons))
}
