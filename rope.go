package ropeExperiment

import (
	"fmt"
	"io"
)

// Rope interface has the functionality for inserting and removing text
// from the rope structure
type Rope interface {
	fmt.Stringer

	// ByteLength returns the number of bytes for all the runes in the rope
	ByteLength() int

	// Insert adds the provided `value` string starting at `start`
	Insert(start int, value string) error

	// Length returns the number of runes in the rope
	Length() int

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
