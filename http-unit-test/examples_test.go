package sum_test

import (
	"fmt"

	sum "github.com/Bash-Clevin/tdd-golang/http-unit-test"
)

/**
You can run sample test to publish on go doc server
To run godoc:
godoc -http=:6060
**/

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5, 6)
	fmt.Println("sum of one to six is", s)

	// Output:
	// sum of one to six is 21
}
