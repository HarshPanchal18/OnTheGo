package main

import "fmt"

func collections() {

	var x [5]float64
	fmt.Printf("x: %v\n", x)

	x[4] = 10
	fmt.Printf("x: %v\n", x)

	println(x[4])

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	println(total / float64(len(x)))

	total = 0

	for _, value := range x {
		total += value
	}

	println(total / float64(len(x)))

	xx := [5]float64{ 98, 93, 77, 82, 83}

	fmt.Printf("x: %v\n",xx)

	/*
	* A slice is a segment of an array. Like arrays slices are indexable and have a length.
	* Unlike arrays this length is allowed to change.
	*/
	var floatArr []float32
	fmt.Printf("floatArr: %v\n",floatArr)

	/*
	* The only difference between this and an array is the missing length between the brackets.
	* In this case arr has been created with a length of 0.
	*/

	// If you want to create a slice you should use the built-in make function:
	array5 := make([]int, 5) // This creates a slice that is associated with an underlying int[] of length 5.
	fmt.Printf("array5: %v\n",array5)

	var arr10 = array5[0:5]
	fmt.Printf("arr10: %v\n",arr10)

	fmt.Printf("xx: %v\n",xx[0:3])
	fmt.Printf("xx: %v\n",xx[1:])


	// Slices

	slice1 := []int {1, 3, 5}
	// append() - creates a new slice by taking an existing slice (the first argument) and appending all the following arguments to it.
	slice2 := append(slice1, 7, 9)
	fmt.Println(slice1, slice2)

	slice3 := make([]int,2)
	copy(slice3, slice1) // copy(dst, src) - copies elements from a source slice into a destination slice.
	fmt.Println(slice1, slice3)

	// Maps
	person := make(map[string]int)
	person["age"] = 10
	person["street"] = 56
	person["grade"] = 49
	fmt.Println(person)

	delete(person, "street") // nothing happened if the key doesn't exist.
	fmt.Println(person)

	name, ok := person["age"]
	fmt.Println(name, ok) // value - true if the key exists otherwise zero(0) - false

	if grade, ok := person["grade"]; ok {
		fmt.Println(grade, ok)
	}

	person1 := map[string]any {
		"name": "Harsh",
		"grade": 5,
	}

	fmt.Println(person1)

	// First we try to get the value from the map, then if it's successful we run the code inside of the block.
	if grade, ok := person1["grade1"]; ok {
		fmt.Println(grade, ok)
	} else {
		fmt.Println("Unknown. Adding to the map...")
		person1["grade1"] = 95
	}

	fmt.Println(person1)

	elements := map[string]map[string]string{
		"H": {
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": {
			"name":"Helium",
			"state":"gas",
		},
		"Li": {
			"name":"Lithium",
			"state":"solid",
		},
		"Be": {
			"name":"Beryllium",
			"state":"solid",
		},
		"B": {
			"name":"Boron",
			"state":"solid",
		},
		"C": {
			"name":"Carbon",
			"state":"solid",
		},
		"N": {
			"name":"Nitrogen",
			"state":"gas",
		},
		"O": {
			"name":"Oxygen",
			"state":"gas",
		},
		"F": {
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne": {
			"name":"Neon",
			"state":"gas",
		}, // Comma is mandatory for ending the map values.
	}

	// fmt.Printf("Elements: %v\n",elements)

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}