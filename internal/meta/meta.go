package meta

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Meta struct {
	Version string
	Commit  string
}

func (m Meta) UserAgent() string {
	return fmt.Sprintf(
		"%s/%s (%s/%s) %s", command(), m.Version, runtime.GOOS, runtime.GOARCH, m.Commit,
	)
}

func command() string {
	a := os.Args[0]

	if len(a) == 0 {
		return "unknown"
	}

	return filepath.Base(a)
}
