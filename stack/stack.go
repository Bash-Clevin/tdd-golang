package stack

type Stack struct {
	size int
}

func NewStack() *Stack {
	return &Stack{size: 0}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Add(val string) {
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}
