package review

import (
	"github.com/jedib0t/go-pretty/table"
)

type Reviews []Review

func (r Reviews) String() string {
	t := table.NewWriter()

	t.AppendHeader(table.Row{
		"Repository",
		"User",
		"Title",
		"Link",
	})

	for _, review := range r {
		t.AppendRow(review.ToRow())
	}

	return t.Render()
}
