# buildkite-mcp-server

This is an [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction) server for [Buildkite](https://buildkite.com). The goal is to provide access to information from buildkite about pipelines, builds and jobs to tools such as [Claude Desktop](https://claude.ai/download), [GitHub Copilot](https://github.com/features/copilot) and other tools, or editors.

# Tools

* `get_pipeline` - Get details of a specific pipeline in Buildkite
* `list_pipelines` - List all pipelines in a buildkite organization
* `list_builds` - List all builds in a pipeline in Buildkite
* `list_pipelines` - List all pipelines in a buildkite organization

Example of the `get_pipeline` tool in action.

![Get Pipeline Tool](docs/images/get_pipeline.png)

# prerequisites

* [goreleaser](http://goreleaser.com)
* [go 1.24](https://go.dev)

# building

Build the binary.

```
make build
```

Copy it to your path.

# configuration

Create a buildkite api token with read access to pipelines.

```json
{
    "mcpServers": {
        "buildkite": {
            "command": "buildkite-mcp-server",
            "args": [
                "stdio"
            ],
            "env": {
                "BUILDKITE_API_TOKEN": "bkua_xxxxxxxx"
            }
        }
    }
}
```


## Disclaimer

This project is in the early stages of development and is not yet ready for use.

## License

This project is released under MIT license.