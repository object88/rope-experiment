package ropeExperiment

import (
	"bytes"
	"fmt"
	"io"
)

type V3 struct {
	right  *V3
	left   *V3
	value  *[]byte
	length int
}

func CreateV3(initial string) *V3 {
	b := []byte(initial)
	r := &V3{nil, nil, &b, len(initial)}
	r.adjust()
	return r
}

func CreateV3FromBytes(initial []byte) *V3 {
	r := &V3{nil, nil, &initial, len(initial)}
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
	read := r.NewReader()
	io.Copy(&buf, read)
	return string(buf.Bytes())
}

func (r *V3) adjust() {
	if r.value != nil {
		if r.length > splitLength {
			divide := r.length >> 1
			r.left = CreateV3FromBytes((*r.value)[:divide])
			r.right = CreateV3FromBytes((*r.value)[divide:])
			r.value = nil
		}
	} else {
		if r.length < joinLength {
			v := r.left.String() + r.right.String()
			b := []byte(v)
			r.value = &b
			r.left = nil
			r.right = nil
		}
	}
}

func (r *V3) insert(position int, value string) error {
	if r.value != nil {
		var buf bytes.Buffer
		buf.Grow(r.length + len(value))
		buf.Write((*r.value)[0:position])
		buf.Write([]byte(value))
		buf.Write((*r.value)[position:])
		b := buf.Bytes()
		r.value = &b
		r.length = len(*r.value)
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
		v := r.left.String() + r.right.String()
		b := []byte(v)
		r.value = &b
		r.left = nil
		r.right = nil
		r.adjust()
	}
}

func (r *V3) remove(start, end int) error {
	if r.value != nil {
		var buf bytes.Buffer
		buf.Grow(len(*r.value) - end + start)
		buf.Write((*r.value)[0:start])
		buf.Write((*r.value)[end:])
		b := buf.Bytes()
		r.value = &b
		r.length = len(*r.value)
	} else {
		leftLength := r.left.length
		leftStart := min(start, leftLength)
		leftEnd := min(end, leftLength)
		rightLength := r.right.length
		rightStart := max(0, min(start-leftLength, rightLength))
		rightEnd := max(0, min(end-leftLength, rightLength))
		if leftStart < leftLength {
			r.left.remove(leftStart, leftEnd)
		}
		if rightEnd > 0 {
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
	if read.pos == read.r.length {
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
		if copied != r.length && err == nil {
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
