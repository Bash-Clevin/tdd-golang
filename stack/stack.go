package stack

type Stack struct {
	size   int
	values []string
}

func NewStack() *Stack {
	return &Stack{size: 0, values: make([]string, 5)}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Add(val string) {
	s.values[s.size] = val
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Pop() string {
	s.size--
	return s.values[s.size]
}
