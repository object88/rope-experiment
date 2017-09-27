package ropeExperiment

import (
	"fmt"
	"io"
)

// Rope interface has the functionality for inserting and removing text
// from the rope structure
type Rope interface {
	fmt.Stringer

	// Insert adds the provided `value` string starting at `start`
	Insert(start int, value string) error

	// NewReader returns an io.Reader over the Rope
	NewReader() io.Reader

	// Remove removes the characters from `start` to `end`
	Remove(start, end int) error
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
