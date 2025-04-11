package buildkite

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/buildkite/go-buildkite/v4"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type UserClient interface {
	CurrentUser(ctx context.Context) (buildkite.User, *buildkite.Response, error)
}

func CurrentUser(ctx context.Context, client UserClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("current_user",
			mcp.WithDescription("Get the current user"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			user, resp, err := client.CurrentUser(ctx)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			if resp.StatusCode != 200 {
				return mcp.NewToolResultError("failed to get current user"), nil
			}

			r, err := json.Marshal(&user)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal user: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}
