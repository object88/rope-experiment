package ropeExperiment

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

type V3 struct {
	right      *V3
	left       *V3
	value      *[]byte
	length     int
	byteLength int
}

func CreateV3(initial string) *V3 {
	b := []byte(initial)
	r := &V3{nil, nil, &b, utf8.RuneCountInString(initial), len(initial)}
	r.adjust()
	return r
}

func CreateV3FromBytes(initial []byte) *V3 {
	r := &V3{nil, nil, &initial, utf8.RuneCount(initial), len(initial)}
	r.adjust()
	return r
}

func (r *V3) Insert(position int, value string) error {
	if r == nil {
		return fmt.Errorf("Nil pointer receiver")
	}

	if position < 0 || position > r.length {
		return fmt.Errorf("position is not within rope bounds")
	}

	return r.insert(position, value)
}

func (r *V3) NewReader() io.Reader {
	return &V3Reader{0, r}
}

// Rebalance rebalances the b-tree structure
func (r *V3) Rebalance() {
	if r.value == nil {
		leftLength := r.left.length
		rightLength := r.right.length

		if float32(leftLength)/float32(rightLength) > rebalanceRatio ||
			float32(rightLength)/float32(leftLength) > rebalanceRatio {
			r.rebuild()
		} else {
			r.left.Rebalance()
			r.right.Rebalance()
		}
	}
}

func (r *V3) Remove(start, end int) error {
	if r == nil {
		return fmt.Errorf("Nil pointer receiver")
	}

	if start < 0 || start > r.length {
		return fmt.Errorf("Start is not within rope bounds")
	}
	if end < 0 || end > r.length {
		return fmt.Errorf("End is not within rope bounds")
	}
	if start > end {
		return fmt.Errorf("Start is greater than end")
	}

	return r.remove(start, end)
}

func (r *V3) String() string {
	var buf bytes.Buffer
	buf.Grow(r.byteLength)
	read := r.NewReader()
	io.Copy(&buf, read)
	return string(buf.Bytes())
}

func (r *V3) adjust() {
	if r.value != nil {
		if r.length > splitLength {
			divide := r.length >> 1
			offset := r.findByteOffsets(divide)
			r.left = CreateV3FromBytes((*r.value)[:offset])
			r.right = CreateV3FromBytes((*r.value)[offset:])
			r.value = nil
		}
	} else {
		if r.length < joinLength {
			r.join()
		}
	}
}

func (r *V3) insert(position int, value string) error {
	if r.value != nil {
		var buf bytes.Buffer
		offset := r.findByteOffsets(position)
		valueLength := utf8.RuneCountInString(value)
		buf.Grow(r.length + len(value))
		buf.Write((*r.value)[0:offset])
		buf.Write([]byte(value))
		buf.Write((*r.value)[offset:])
		b := buf.Bytes()
		r.value = &b
		r.length += valueLength
	} else {
		leftLength := r.left.length
		if position < leftLength {
			r.left.insert(position, value)
			r.length = r.left.length + r.right.length
		} else {
			r.right.insert(position-leftLength, value)
		}
	}
	r.adjust()
	return nil
}

func (r *V3) join() {
	c := r.left.byteLength + r.right.byteLength
	var buf bytes.Buffer
	buf.Grow(c)
	io.Copy(&buf, r.left.NewReader())
	io.Copy(&buf, r.right.NewReader())
	b := buf.Bytes()
	r.value = &b
	r.left = nil
	r.right = nil
}

func (r *V3) findByteOffsets(position int) int {
	rs := []rune(string(*r.value))

	offset := 0

	for i := 0; i < position; i++ {
		offset += utf8.RuneLen(rs[i])
	}

	return offset
}

func (r *V3) locate(position int) (*V3, int) {
	if r.value != nil {
		return r, position
	}

	leftLength := r.left.length
	if position < leftLength {
		return r.left.locate(position)
	}

	return r.right.locate(position - leftLength)
}

func (r *V3) rebuild() {
	if r.value == nil {
		r.join()
		r.adjust()
	}
}

func (r *V3) remove(start, end int) error {
	if r.value != nil {
		var buf bytes.Buffer
		byteStart := r.findByteOffsets(start)
		byteEnd := r.findByteOffsets(end)
		buf.Grow(len(*r.value) - byteEnd + byteStart)
		buf.Write((*r.value)[0:byteStart])
		buf.Write((*r.value)[byteEnd:])
		b := buf.Bytes()
		r.value = &b
		r.length -= end - start
	} else {
		leftLength := r.left.length
		leftStart := min(start, leftLength)
		rightLength := r.right.length
		rightEnd := max(0, min(end-leftLength, rightLength))
		if leftStart < leftLength {
			leftEnd := min(end, leftLength)
			r.left.remove(leftStart, leftEnd)
		}
		if rightEnd > 0 {
			rightStart := max(0, min(start-leftLength, rightLength))
			r.right.remove(rightStart, rightEnd)
		}
		r.length = r.left.length + r.right.length
	}

	r.adjust()
	return nil
}

type V3Reader struct {
	pos int
	r   *V3
}

func (read *V3Reader) Read(p []byte) (n int, err error) {
	// NOTE: method not currently tested, because test code is invoking WriteTo
	// instead.
	if read.pos == read.r.byteLength {
		return 0, io.EOF
	}

	node, offset := read.r.locate(read.pos)

	copied := copy(p, (*node.value)[offset:])
	read.pos += copied
	return copied, nil
}

func (read *V3Reader) WriteTo(w io.Writer) (int64, error) {
	n, err := read.writeNodeTo(read.r, w)
	return int64(n), err
}

func (read *V3Reader) writeNodeTo(r *V3, w io.Writer) (int, error) {
	if r.value != nil {
		copied, err := w.Write(*r.value)
		if copied != len(*r.value) && err == nil {
			err = io.ErrShortWrite
		}
		return copied, err
	}

	var err error
	var n, m int

	n, err = read.writeNodeTo(r.left, w)
	if err != nil {
		return n, err
	}

	m, err = read.writeNodeTo(r.right, w)

	return int(n + m), err
}
