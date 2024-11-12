package main

type testPair struct {
	values []float64
	avg    float64
}

var tests = []testPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1}, 2},
	{[]float64{-1, 1}, 0},
}
