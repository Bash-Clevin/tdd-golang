package sum_test

import (
	"testing"

	sum "github.com/Bash-Clevin/tdd-golang/http-unit-test"
)

/*
*
This is a table driven test
Ref: https://go.dev/blog/subtests
*
*/
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

/** what if you have multiple tests
	and you also want to choose the test to run by name
	i.e
	go test -v -run <name of test>

	go test -v -run Ints/one_to_five

	Note: name of test is regex so it can take in . or match
**/

func TestSubInts(t *testing.T) {

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
		t.Run(testCase.name, func(t *testing.T) {
			s := sum.Ints(testCase.numbers...)

			if s != testCase.sum {
				t.Fatalf("Sum of %v should be %d but got %d", testCase.name, testCase.sum, s)
			}
		})

	}
}
