package ropeExperiment

import (
	"fmt"
	"io"
	"unicode/utf8"
)

// This code is a mostly-direct translation of
// https://github.com/component/rope.  Many thanks to the contributers and
// maintainers of http://component.github.io/ for their unknown contributions
// to this project.

const (
	splitLength = 512
	joinLength  = 256

	rebalanceRatio = 1.2
)

type V2 struct {
	right  *V2
	left   *V2
	value  *string
	length int
}

func CreateV2(initial string) *V2 {
	r := &V2{nil, nil, &initial, utf8.RuneCountInString(initial)}
	r.adjust()
	return r
}

func (r *V2) Insert(position int, value string) error {
	if r == nil {
		return fmt.Errorf("Nil pointer receiver")
	}

	if position < 0 || position > r.length {
		return fmt.Errorf("position is not within rope bounds")
	}

	return r.insert(position, value)
}

func (r *V2) NewReader() io.Reader {
	return &V2Reader{0, r}
}

// Rebalance rebalances the b-tree structure
func (r *V2) Rebalance() {
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

func (r *V2) Remove(start, end int) error {
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

func (r *V2) String() string {
	if r.value != nil {
		return *r.value
	}
	return r.left.String() + r.right.String()
}

func (r *V2) adjust() {
	if r.value != nil {
		if r.length > splitLength {
			divide := r.length >> 1
			r.left = CreateV2(string([]rune(*r.value)[:divide]))
			r.right = CreateV2(string([]rune(*r.value)[divide:]))
			r.value = nil
		}
	} else {
		if r.length < joinLength {
			v := r.left.String() + r.right.String()
			r.value = &v
			r.left = nil
			r.right = nil
		}
	}
}

func (r *V2) insert(position int, value string) error {
	if r.value != nil {
		v := string([]rune(*r.value)[0:position]) + value + string([]rune(*r.value)[position:])
		r.value = &v
		r.length = utf8.RuneCountInString(*r.value)
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

func (r *V2) locate(position int) (*V2, int) {
	if r.value != nil {
		return r, position
	}

	leftLength := r.left.length
	if position < leftLength {
		return r.left.locate(position)
	}

	return r.right.locate(position - leftLength)
}

func (r *V2) rebuild() {
	if r.value == nil {
		v := r.left.String() + r.right.String()
		r.value = &v
		r.left = nil
		r.right = nil
		r.adjust()
	}
}

func (r *V2) remove(start, end int) error {
	if r.value != nil {
		v := string([]rune(*r.value)[0:start]) + string([]rune(*r.value)[end:])
		r.value = &v
		r.length = utf8.RuneCountInString(*r.value)
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

type V2Reader struct {
	pos int
	r   *V2
}

func (read *V2Reader) Read(p []byte) (n int, err error) {
	// NOTE: method not currently tested, because test code is invoking WriteTo
	// instead.
	if read.pos == read.r.length {
		return 0, io.EOF
	}

	node, offset := read.r.locate(read.pos)

	copied := copy(p, []byte(string([]rune(*node.value)[offset:])))
	read.pos += copied
	return copied, nil
}

func (read *V2Reader) WriteTo(w io.Writer) (int64, error) {
	n, err := read.writeNodeTo(read.r, w)
	return int64(n), err
}

func (read *V2Reader) writeNodeTo(r *V2, w io.Writer) (int, error) {
	if r.value != nil {
		copied, err := io.WriteString(w, *r.value)
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
