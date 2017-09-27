package ropeExperiment

import (
	"bytes"
	"io"
	"strings"
	"unicode/utf8"
)

// V1 is a simple string
type V1 struct {
	value string
}

func CreateV1(initial string) Rope {
	return &V1{initial}
}

func (r *V1) ByteLength() int {
	return len(r.value)
}

func (r *V1) Insert(start int, value string) error {
	var buf bytes.Buffer
	offset := r.findByteOffsets(start)
	buf.Grow(len(r.value) + len(value))
	buf.WriteString(r.value[:offset])
	buf.WriteString(value)
	buf.WriteString(r.value[offset:])
	r.value = buf.String()
	return nil
}

func (r *V1) Length() int {
	return utf8.RuneCountInString(r.value)
}

func (r *V1) NewReader() io.Reader {
	return strings.NewReader(r.value)
}

func (r *V1) Remove(start, end int) error {
	var buf bytes.Buffer
	byteStart := r.findByteOffsets(start)
	byteEnd := r.findByteOffsets(end)
	buf.Grow(len(r.value) - byteEnd + byteStart)
	buf.WriteString(r.value[:byteStart])
	buf.WriteString(r.value[byteEnd:])
	r.value = buf.String()
	return nil
}

func (r *V1) String() string {
	return r.value
}

func (r *V1) findByteOffsets(position int) int {
	rs := []rune(r.value)

	offset := 0

	for i := 0; i < position; i++ {
		offset += utf8.RuneLen(rs[i])
	}

	return offset
}
