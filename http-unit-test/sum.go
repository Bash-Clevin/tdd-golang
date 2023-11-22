package sum

// Int returns the sum of a list of integers

func Ints(values ...int) int {
	return ints(values)
}

func ints(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return ints(values[1:]) + values[0]
}
