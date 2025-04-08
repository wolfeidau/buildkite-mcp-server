package main

import (
	"context"
	"log"

	"github.com/alecthomas/kong"
	"github.com/buildkite/go-buildkite/v4"
	"github.com/wolfeidau/buildkite-mcp-server/internal/commands"
)

var (
	version = "dev"

	cli struct {
		Stdio    commands.StdioCmd `cmd:"" help:"stdio mcp server."`
		APIToken string            `help:"The Buildkite API token to use." env:"BUILDKITE_API_TOKEN"`
		Debug    bool              `help:"Enable debug mode."`
		Version  kong.VersionFlag
	}
)

func main() {
	ctx := context.Background()
	cmd := kong.Parse(&cli,
		kong.Name("buildkite-mcp-server"),
		kong.Description("A server that proxies requests to the Buildkite API."),
		kong.UsageOnError(),
		kong.Vars{
			"version": version,
		},
		kong.BindTo(ctx, (*context.Context)(nil)),
	)

	client, err := buildkite.NewOpts(buildkite.WithTokenAuth(cli.APIToken))
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run(&commands.Globals{Debug: cli.Debug, Version: version, Client: client})
	cmd.FatalIfErrorf(err)
}
