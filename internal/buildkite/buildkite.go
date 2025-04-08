package buildkite

import (
	"errors"

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
