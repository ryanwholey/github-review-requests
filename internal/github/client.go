package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/ryanwholey/github-review-requests/internal/meta"
	"golang.org/x/oauth2"
)

type Client struct {
	username string
	client   *github.Client
}

func New(ctx context.Context, username string, token string, meta meta.Meta) Client {
	client := github.NewClient(
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})),
	)

	client.UserAgent = meta.UserAgent()

	return Client{
		client:   client,
		username: username,
	}
}

func (c Client) GetReviews(ctx context.Context) (*github.IssuesSearchResult, *github.Response, error) {
	return c.client.Search.Issues(
		ctx,
		fmt.Sprintf("is:open is:pr review-requested:%s archived:false", c.username),
		&github.SearchOptions{},
	)
}
