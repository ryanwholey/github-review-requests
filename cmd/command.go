package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/ryanwholey/github-review-requests/internal/github"
	"github.com/ryanwholey/github-review-requests/internal/io"
	"github.com/ryanwholey/github-review-requests/internal/run"
	"github.com/ryanwholey/github-review-requests/internal/storage"
	"github.com/spf13/cobra"
)

func NewCommand(streams io.Streams) *cobra.Command {
	var username string
	var storagePath string
	var interval time.Duration
	var clean bool

	cmd := &cobra.Command{
		Use: "github-review-requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			token := os.Getenv("GH_TOKEN")
			if token == "" {
				return fmt.Errorf("environment variable GH_TOKEN must be set")
			}

			gh := github.New(ctx, username, token)
			sm := storage.NewManager(storagePath, clean)

			if interval == 0 {
				return run.Run(ctx, gh, sm, streams)
			}

			// TODO: Implement timer loop

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&username, "username", "u", os.Getenv("GH_USERNAME"), "GitHub username")
	flags.StringVarP(&storagePath, "storage", "s", "~/.github-review-request-storage", "GitHub review storage")
	flags.DurationVarP(&interval, "interval", "i", 0, "Run the command on an interval")
	flags.BoolVarP(&clean, "clean", "c", false, "Whether to remove storage before execution")

	return cmd
}
