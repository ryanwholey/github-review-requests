package main

import (
	"os"

	"github.com/ryanwholey/github-review-requests/cmd"
	"github.com/ryanwholey/github-review-requests/internal/io"
	"github.com/ryanwholey/github-review-requests/internal/meta"
	"github.com/spf13/cobra"
)

var version string = "dev"
var commit string = "0000000"

func main() {
	command := cmd.NewCommand(meta.Meta{
		Version: version,
		Commit:  commit,
	}, io.Streams{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	cobra.CheckErr(command.Execute())
}
