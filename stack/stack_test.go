package stack_test

import (
	"testing"

	"github.com/tdd-golang/stack"
)

func TestEmpty(t *testing.T) {
	s := stack.NewStack()

	if s.Empty() != true {
		t.Error("Stack was not empty")
	}
}

func TestNotEmpty(t *testing.T) {
	s := stack.NewStack()
	s.Add("Bob")

	if s.Empty() != false {
		t.Error("Stack was Empty")
	}
}

func TestSizeZero(t *testing.T) {
	s := stack.NewStack()

	if s.Size() != 0 {
		t.Errorf("Expected zero elements, found: %d", s.Size())
	}
}

func TestSizeOne(t *testing.T) {
	s := stack.NewStack()
	s.Add("Bob")

	if s.Size() != 1 {
		t.Error("Incorrect size")
		t.Log("Expected: 1")
		t.Logf("Actual: %d", s.Size())
	}
}

func TestSizeThree(t *testing.T) {
	s := stack.NewStack()
	s.Add("Bob")
	s.Add("Alex")
	s.Add("Miriam")

	if s.Size() != 3 {
		t.Errorf("Expected 3, found %d", s.Size())
	}
}
