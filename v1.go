package ropeExperiment

import (
	"io"
	"strings"
)

// V1 is a simple string
type V1 struct {
	value string
}

func CreateV1(initial string) Rope {
	return &V1{initial}
}

func (r *V1) Insert(start int, value string) error {
	r.value = r.value[0:start] + value + r.value[start:]
	return nil
}

func (r *V1) NewReader() io.Reader {
	return strings.NewReader(r.value)
}

func (r *V1) Remove(start, end int) error {
	r.value = r.value[0:start] + r.value[end:]
	return nil
}
