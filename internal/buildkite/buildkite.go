package buildkite

import (
	"errors"

	"github.com/buildkite/go-buildkite/v4"
	"github.com/mark3labs/mcp-go/mcp"
)

// requiredParam validates and retrieves a required parameter of a specified type from a CallToolRequest.
// It checks for the parameter's presence, type correctness, and non-emptiness.
// Returns the parameter value if valid, or an error describing the validation failure.
func requiredParam[T comparable](request mcp.CallToolRequest, name string) (T, error) {
	var empty T

	if request.Params.Arguments[name] == nil {
		return empty, errors.New("missing required argument: " + name)
	}

	val, ok := request.Params.Arguments[name].(T)
	if !ok {
		return empty, errors.New("invalid type for argument: " + name)
	}

	if val == empty {
		return empty, errors.New("argument cannot be empty: " + name)
	}

	return val, nil
}

func optionalPaginationParams(r mcp.CallToolRequest) (buildkite.ListOptions, error) {
	page, err := optionalParamWithDefault(r, "page", 1)
	if err != nil {
		return buildkite.ListOptions{}, err
	}
	perPage, err := optionalParamWithDefault(r, "perPage", 30)
	if err != nil {
		return buildkite.ListOptions{}, err
	}
	return buildkite.ListOptions{
		Page:    page,
		PerPage: perPage,
	}, nil
}

func optionalParamWithDefault[T comparable](r mcp.CallToolRequest, name string, defaultValue T) (T, error) {
	if r.Params.Arguments[name] == nil {
		return defaultValue, nil
	}

	val, ok := r.Params.Arguments[name].(T)
	if !ok {
		return defaultValue, errors.New("invalid type for argument: " + name)
	}

	return val, nil
}

func withPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("page",
			mcp.Description("Page number for pagination (min 1)"),
			mcp.Min(1),
		)(tool)

		mcp.WithNumber("perPage",
			mcp.Description("Results per page for pagination (min 1, max 100)"),
			mcp.Min(1),
			mcp.Max(100),
		)(tool)
	}
}
