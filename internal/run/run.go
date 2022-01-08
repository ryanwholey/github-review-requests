package run

import (
	"context"
	"fmt"

	"github.com/ryanwholey/github-review-requests/internal/github"
	"github.com/ryanwholey/github-review-requests/internal/io"
	"github.com/ryanwholey/github-review-requests/internal/notify"
	"github.com/ryanwholey/github-review-requests/internal/review"
	"github.com/ryanwholey/github-review-requests/internal/storage"
)

func Run(ctx context.Context, gh github.Client, sm storage.Manager, streams io.Streams) error {
	store, err := sm.Load()
	if err != nil {
		return err
	}

	result, _, err := gh.GetReviews(ctx)
	if err != nil {
		return err
	}

	reviews := make(review.Reviews, result.GetTotal())

	for i, issue := range result.Issues {
		reviews[i] = review.New(issue)
	}

	added := store.FindNew(reviews)

	notifier := notify.New()

	for _, rev := range added {
		notifier.Notify(rev.ToNote())
	}

	if err := sm.Save(storage.Storage{Reviews: reviews}); err != nil {
		return err
	}

	fmt.Fprintf(streams.Stdout, "%v\n", reviews)

	return nil
}
