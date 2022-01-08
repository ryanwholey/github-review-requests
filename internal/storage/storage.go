package storage

import "github.com/ryanwholey/github-review-requests/internal/review"

type Storage struct {
	Reviews []review.Review `json:"reviews"`
}

func (s Storage) FindNew(updates []review.Review) []review.Review {
	added := []review.Review{}

	for _, u := range updates {
		isNew := true

		for _, r := range s.Reviews {
			if u.ID == r.ID {
				isNew = false
			}
		}

		if isNew {
			added = append(added, u)
		}
	}

	return added
}
