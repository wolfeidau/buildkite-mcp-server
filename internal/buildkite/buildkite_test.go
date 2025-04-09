package buildkite

import (
	"testing"

	"github.com/buildkite/go-buildkite/v4"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/require"
)

func Test_requiredParam(t *testing.T) {
	t.Run("valid string parameter", func(t *testing.T) {
		assert := require.New(t)
		req := createMCPRequest(t, map[string]any{
			"test": "value",
		})

		result, err := requiredParam[string](req, "test")
		assert.NoError(err)
		assert.Equal("value", result)
	})
}

func Test_optionalPaginationParams(t *testing.T) {
	assert := require.New(t)
	req := createMCPRequest(t, map[string]any{
		"page":    1,
		"perPage": 30,
	})

	opts, err := optionalPaginationParams(req)
	assert.NoError(err)
	assert.Equal(buildkite.ListOptions{
		Page:    1,
		PerPage: 30,
	}, opts)
}

func createMCPRequest(t *testing.T, args map[string]any) mcp.CallToolRequest {
	t.Helper()
	return mcp.CallToolRequest{
		Params: struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments,omitempty"`
			Meta      *struct {
				ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
			} `json:"_meta,omitempty"`
		}{
			Arguments: args,
		},
	}
}
