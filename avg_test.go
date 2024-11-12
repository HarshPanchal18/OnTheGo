package main

import "testing"

func testAverage(t *testing.T) {
	v := Average([]float64{1, 2})

	if v != 1.3 {
		t.Errorf("expected 1.5, got %v", v)
	} else {
		t.Log("Passed...")
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
*/