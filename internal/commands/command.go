package commands

import (
	"github.com/buildkite/go-buildkite/v4"
)

type Globals struct {
	Client  *buildkite.Client
	Version string
	Debug   bool
}
