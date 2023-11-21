package stack

type Stack struct {
	isEmpty bool
	size    int
}

func NewStack() *Stack {
	return &Stack{isEmpty: true, size: 0}
}

func (s *Stack) Empty() bool {
	return s.isEmpty
}

func (s *Stack) Add(val string) {
	s.isEmpty = false
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}
