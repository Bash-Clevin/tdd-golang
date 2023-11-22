package sum_test

import (
	"testing"

	sum "github.com/Bash-Clevin/tdd-golang/http-unit-test"
)

func TestInts(t *testing.T) {
	s := sum.Ints(1, 2, 3, 4, 5, 7)

	if s != 15 {
		t.Errorf("Expected a sum of 15 but got %d", s)
	}
}
