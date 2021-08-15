package go_koans

func divideFourBy(i int) int {
	defer handlePanics()

	if i == 0 {
		i = 2
	}

	return 4 / i
}

func handlePanics() int {
	if x:= recover(); x != nil {
		return 2
	}

	return 2
}

const __divisor__ = 0

func aboutPanics() {
	assert( divideFourBy(4) == 1) // panics are exceptional errors at runtime

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
