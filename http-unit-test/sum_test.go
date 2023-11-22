package sum_test

import (
	"testing"

	sum "github.com/Bash-Clevin/tdd-golang/http-unit-test"
)

func TestInts(t *testing.T) {

	testingTable := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{name: "one to five", numbers: []int{1, 2, 3, 4, 5}, sum: 15},
		{name: "nil", numbers: nil, sum: 0},
		{name: "one and minus one", numbers: []int{1, -1}, sum: 0},
	}

	for _, testCase := range testingTable {
		s := sum.Ints(testCase.numbers...)

		if s != testCase.sum {
			t.Errorf("Sum of %v should be %d but got %d", testCase.name, testCase.sum, s)
		}
	}
}
