package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ryanwholey/github-review-requests/internal/github"
	"github.com/ryanwholey/github-review-requests/internal/io"
	"github.com/ryanwholey/github-review-requests/internal/meta"
	"github.com/ryanwholey/github-review-requests/internal/run"
	"github.com/ryanwholey/github-review-requests/internal/storage"
	"github.com/spf13/cobra"
)

func NewCommand(meta meta.Meta, streams io.Streams) *cobra.Command {
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

			gh := github.New(ctx, username, token, meta)
			sm := storage.NewManager(storagePath, clean)

			first := run.Run(ctx, gh, sm, streams)

			if interval == 0 {
				return first
			}

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt)

			timer := time.Tick(interval)

			for {
				select {
				case <-timer:
					if err := run.Run(ctx, gh, sm, streams); err != nil {
						return err
					}
				case <-sigChan:
					return nil
				}
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&username, "username", "u", os.Getenv("GH_USERNAME"), "GitHub username")
	flags.StringVarP(&storagePath, "storage", "s", "~/.github-review-request-storage", "GitHub review storage")
	flags.DurationVarP(&interval, "interval", "i", 0, "Run the command on an interval")
	flags.BoolVarP(&clean, "clean", "c", false, "Whether to remove storage before execution")

	return cmd
}
