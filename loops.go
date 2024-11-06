package main

import "fmt"

func loops() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for i := 4; i > 0; i-- {
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		if(i % 2 == 0) {
			fmt.Println(i, " Even")
		} else {
			fmt.Println(i, " Odd")
		}
	}

	for i := 1; i <= 5; i++ {
		switch i {

		case 1: fmt.Println("one")
		case 2: fmt.Println("two")
		case 3: fmt.Println("three")
		case 4: fmt.Println("four")
		// case 5: fmt.Println("five")
		default: fmt.Println(i)

		}
	}

}