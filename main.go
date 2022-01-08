package main

import (
	"os"

	"github.com/ryanwholey/github-review-requests/cmd"
	"github.com/ryanwholey/github-review-requests/internal/io"
	"github.com/spf13/cobra"
)

func main() {

	command := cmd.NewCommand(io.Streams{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	cobra.CheckErr(command.Execute())
}
