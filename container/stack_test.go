package containert

import (
	"errors"
	"testing"
)

type Stack []interface{}

func New(cap int) Stack {
	return make([]interface{}, 0, cap)
}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s Stack) Top() (interface{}, error) {
	l := len(s)
	if l == 0 {
		return nil, errors.New("out of stack, len is 0")
	}
	return s[l-1], nil
}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (interface{}, error) {
	l := len(*s)
	if l == 0 {
		return nil, errors.New("stack is empty")
	}
	var v interface{}
	v = (*s)[l-1]
	*s = (*s)[:l-1]
	return v, nil
}

func TestStack(t *testing.T) {
	stack := New(3)

	t.Log(stack.Len())
	t.Log(stack.Cap())
	t.Log(stack.IsEmpty())

	stack.Push(2)
	t.Log(stack.IsEmpty())

	t.Log(stack.Top())
	t.Log(stack.Top())

	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	t.Log(stack.Cap())
	t.Log(stack.Pop())
	t.Log(stack.Len())

}
