package ropeExperiment

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

type V4 struct {
	left       *V4
	right      *V4
	value      *[]byte
	length     int
	byteLength int
}

func CreateV4(initial string) *V4 {
	byteLength := len(initial)
	charLength := utf8.RuneCountInString(initial)
	var r *V4
	if len(initial) > splitLength {
		divide := charLength >> 1
		offset := findByteOffset(initial, divide)
		r = &V4{
			CreateV4(initial[:offset]),
			CreateV4(initial[offset:]),
			nil,
			charLength,
			byteLength,
		}
	} else {
		b := make([]byte, splitLength)
		copy(b, initial)
		r = &V4{
			nil,
			nil,
			&b,
			charLength,
			byteLength,
		}
	}

	return r
}

func CreateV4FromBytes(initial []byte) *V4 {
	r := &V4{nil, nil, &initial, utf8.RuneCount(initial), len(initial)}
	r.adjust()
	return r
}

func createChildrenFromBytes(f ...[]byte) (*V4, *V4) {
	totalLength := 0
	totalByteLength := 0
	for _, s := range f {
		for i := 0; i < len(s); {
			n := utf8.RuneCount(s)
			totalLength += n
			totalByteLength += len(s)
		}
	}

	// divide := totalLength >> 1

	return nil, nil
}

func createTree(length, byteLength int) *V4 {

	return nil
}

func (r *V4) ByteLength() int {
	return r.byteLength
}

func (r *V4) Insert(position int, value string) error {
	if r == nil {
		return fmt.Errorf("Nil pointer receiver")
	}

	if position < 0 || position > r.length {
		return fmt.Errorf("position is not within rope bounds")
	}

	return r.insert(position, value)
}

func (r *V4) Length() int {
	return r.length
}

func (r *V4) NewReader() io.Reader {
	return &V4Reader{0, r}
}

// Rebalance rebalances the b-tree structure
func (r *V4) Rebalance() {
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

func (r *V4) Remove(start, end int) error {
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

func (r *V4) String() string {
	read := r.NewReader()
	b := make([]byte, r.byteLength)
	for n := 0; n < r.byteLength; {
		m, _ := read.Read(b[n:])
		n += m
	}
	return string(b)
}

func (r *V4) adjust() {
	if r.value != nil {
		if r.length > splitLength {
			divide := r.length >> 1
			offset := r.findByteOffsets(divide)
			r.left = CreateV4FromBytes((*r.value)[:offset])
			r.right = CreateV4FromBytes((*r.value)[offset:r.byteLength])
			r.value = nil
		}
	} else {
		if r.length < joinLength {
			r.join()
		}
	}
}

func (r *V4) insert(position int, value string) error {
	if r.value != nil {
		offset := r.findByteOffsets(position)
		valueLength := utf8.RuneCountInString(value)
		valueBytesLength := len(value)
		if r.byteLength+valueBytesLength >= splitLength {
			copy((*r.value)[:offset+valueBytesLength], (*r.value)[:offset])
			copy((*r.value)[:offset], value)
			r.byteLength += valueBytesLength
			r.length += valueLength
		} else {
			// divide := offset >> 1
		}
		// var buf bytes.Buffer
		// buf.Grow(r.byteLength + valueBytesLength)
		// buf.Write((*r.value)[0:offset])
		// buf.Write([]byte(value))
		// buf.Write((*r.value)[offset:])
		// b := buf.Bytes()
		// r.value = &b
		// r.byteLength += valueBytesLength
		// r.length += valueLength
	} else {
		leftLength := r.left.length
		if position < leftLength {
			r.left.insert(position, value)
		} else {
			r.right.insert(position-leftLength, value)
		}
		r.byteLength = r.left.byteLength + r.right.byteLength
		r.length = r.left.length + r.right.length
	}
	r.adjust()
	return nil
}

func (r *V4) join() {
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

func (r *V4) findByteOffsets(position int) int {
	offset := 0
	rs := []rune(string(*r.value))

	for i := 0; i < position; i++ {
		offset += utf8.RuneLen(rs[i])
	}

	return offset
}

func (r *V4) locate(offset int) (*V4, int) {
	if r.value != nil {
		return r, offset
	}

	leftByteLength := r.left.byteLength
	if offset < leftByteLength {
		return r.left.locate(offset)
	}

	return r.right.locate(offset - leftByteLength)
}

func (r *V4) rebuild() {
	if r.value == nil {
		r.join()
		r.adjust()
	}
}

func (r *V4) remove(start, end int) error {
	if r.value != nil {
		byteStart := r.findByteOffsets(start)
		byteEnd := r.findByteOffsets(end)
		copy((*r.value)[byteStart:], (*r.value)[byteEnd:])
		r.byteLength -= byteEnd - byteStart
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
		r.byteLength = r.left.byteLength + r.right.byteLength
		r.length = r.left.length + r.right.length
	}

	r.adjust()
	return nil
}

// V4Reader implements io.Reader and io.WriterTo for a V4 rope
type V4Reader struct {
	offset int
	r      *V4
}

func (read *V4Reader) Read(p []byte) (n int, err error) {
	if read.offset == read.r.byteLength {
		return 0, io.EOF
	}

	node, offset := read.r.locate(read.offset)

	copied := copy(p, (*node.value)[offset:node.byteLength])
	read.offset += copied
	return copied, nil
}

// WriteTo writes the contents of a V2 to the provided io.Writer
func (read *V4Reader) WriteTo(w io.Writer) (int64, error) {
	n, err := read.writeNodeTo(read.r, w)
	return int64(n), err
}

func (read *V4Reader) writeNodeTo(r *V4, w io.Writer) (int, error) {
	if r.value != nil {
		copied, err := w.Write((*r.value)[:r.byteLength])
		if copied != r.byteLength && err == nil {
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
