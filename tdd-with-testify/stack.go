package tddwithtestify

type Stack struct {
	size int
}

func NewStack() *Stack {
	return &Stack{size: 0}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Bury(item string) {
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}
