package main

import "testing"

func TestAverage(t *testing.T) {

	for _, pair := range tests {
		v := Average(pair.values)

		if v != pair.avg {
			t.Error(
				"For", pair.values,
				"expected", pair.avg,
				"got", v,
			)
		} else {
			t.Logf("Passed for pair %v.", pair.values)
		}
	}
}

func Average(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

/*
The `go test` command will look for any tests in any of the files in the current folder and run them.
Tests are identified by starting a function with the word Test and taking one argument of type *testing.T.
The file should end with _test.go. You've got that right with avg_test.go.
The Go testing framework will automatically find all the functions in your test files that start with Test (like TestAverage) and execute them.
*/