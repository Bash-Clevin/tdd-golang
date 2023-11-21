package tddwithtestify_test

import (
	"testing"

	stack "github.com/Bash-Clevin/tdd-golang/tdd-with-testify"
	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

func (s *StackSuite) TestEmpty() {
	stack := stack.NewStack()

	s.True(stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	stack := stack.NewStack()
	stack.Bury("red")

	s.False(stack.IsEmpty())
}

func (s *StackSuite) TestEmptySizeZero() {
	stack := stack.NewStack()

	s.Zero(stack.Size())
}

func (s *StackSuite) TestTwoSize() {
	stack := stack.NewStack()
	stack.Bury("red")
	stack.Bury("blue")

	s.Equal(2, stack.Size())
}
