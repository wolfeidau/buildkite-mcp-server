package commands

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	"github.com/wolfeidau/buildkite-mcp-server/internal/buildkite"
)

type StdioCmd struct {
}

func (c *StdioCmd) Run(ctx context.Context, globals *Globals) error {

	s := server.NewMCPServer(
		"github-mcp-server",
		globals.Version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging())

	s.AddTool(buildkite.GetPipeline(ctx, globals.Client))
	s.AddTool(buildkite.ListPipeline(ctx, globals.Client))
	s.AddTool(buildkite.ListBuilds(ctx, globals.Client))
	s.AddTool(buildkite.GetBuild(ctx, globals.Client))

	return server.ServeStdio(s)
}
