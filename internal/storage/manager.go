package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryanwholey/github-review-requests/internal/review"
)

type Manager struct {
	path string
}

func NewManager(path string, clean bool) Manager {
	if strings.HasPrefix(path, "~/") {
		dirname, _ := os.UserHomeDir()
		path = filepath.Join(dirname, path[2:])
	}

	m := Manager{path: path}

	if clean {
		m.Delete()
	}

	return m
}

func (m Manager) exists() bool {
	_, err := os.Stat(m.path)

	return err == nil
}

func (m Manager) Load() (Storage, error) {
	var storage Storage

	if !m.exists() {
		return Storage{
			Reviews: []review.Review{},
		}, nil
	}

	data, err := os.ReadFile(m.path)
	if err != nil {
		return storage, err
	}

	if err := json.Unmarshal(data, &storage); err != nil {
		return storage, err
	}

	return storage, nil
}

func (m Manager) Save(storage Storage) error {
	file, err := os.Create(m.path)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}()

	if err != nil {
		return fmt.Errorf("failed to save storage file: %w", err)
	}

	b, err := json.Marshal(storage)
	if err != nil {
		return fmt.Errorf("failed to marshal storage on save: %w", err)
	}

	_, err = file.Write(b)

	return err
}

func (m Manager) Delete() error {
	if m.exists() {
		return os.Remove(m.path)
	}

	return nil
}
