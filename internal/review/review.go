package review

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/jedib0t/go-pretty/table"
	"github.com/ryanwholey/github-review-requests/internal/notify"
)

type Review struct {
	Repository string `json:"review"`
	User       string `json:"user"`
	Title      string `json:"title"`
	Link       string `json:"link"`
	ID         int64  `json:"id"`
}

func (r Review) String() string {
	return fmt.Sprintf("%s:%s [%s]", r.Repository, r.User, r.Title)
}

func New(issue github.Issue) Review {
	var prOwnerName string

	prOwner := issue.GetUser()
	if prOwner != nil {
		prOwnerName = prOwner.GetLogin()
	}

	ru := strings.Split(issue.GetRepositoryURL(), "/")
	repo := strings.Join(ru[len(ru)-2:], "/")

	return Review{
		Repository: repo,
		User:       prOwnerName,
		Title:      issue.GetTitle(),
		Link:       issue.GetHTMLURL(),
		ID:         issue.GetID(),
	}
}

func (r Review) ToRow() table.Row {
	return table.Row{
		r.Repository,
		r.User,
		r.Title,
		r.Link,
	}
}

func (r Review) ToNote() notify.Note {
	return notify.Note{
		Title: fmt.Sprintf("%s - %s", r.User, r.Repository),
		Body:  r.Title,
		Link:  r.Link,
	}
}
