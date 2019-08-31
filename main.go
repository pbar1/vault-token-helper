package main

import (
	"github.com/pbar1/vault-token-helper/internal/command"
)

var (
	Version = ""
	GitCommit = ""
	BuildDate = ""
)

func main() {
	command.Execute(Version, GitCommit, BuildDate)
}
