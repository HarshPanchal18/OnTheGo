package main

import (
	"fmt"
	"sort"

)

type Person struct {
	Name string
	Age  int
}

type ByName []Person
type ByAge []Person

func (a ByName) Len() int           { return len(a) } // Total length
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] } // Swapper
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name } // Comparator

func (a ByAge) Len() int           { return len(a) } // Total length
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] } // Swapper
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } // Comparator

func sortedKids() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}