package ropeExperiment

import "fmt"

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
	r := &V2{nil, nil, &initial, len(initial)}
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
			r.left = CreateV2((*r.value)[:divide])
			r.right = CreateV2((*r.value)[divide:])
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
		v := (*r.value)[0:position] + value + (*r.value)[position:]
		r.value = &v
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
		v := (*r.value)[0:start] + (*r.value)[end:]
		r.value = &v
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

func (r *V2) substring(start, end int) string {
	if end == -1 {
		end = r.length
	}
	if start < 0 {
		start = 0
	} else if start > r.length {
		start = r.length
	}
	if end < 0 {
		end = 0
	} else if end > r.length {
		end = r.length
	}

	if r.value != nil {
		return (*r.value)[start:end]
	}

	leftLength := r.left.length
	leftStart := min(start, leftLength)
	leftEnd := min(end, leftLength)
	rightLength := r.right.length
	rightStart := max(0, min(start-leftLength, rightLength))
	rightEnd := max(0, min(end-leftLength, rightLength))

	if leftStart != leftEnd {
		if rightStart != rightEnd {
			return r.left.substring(leftStart, leftEnd) + r.right.substring(rightStart, rightEnd)
		}

		return r.left.substring(leftStart, leftEnd)
	}

	if rightStart != rightEnd {
		return r.right.substring(rightStart, rightEnd)
	}

	return ""
}
