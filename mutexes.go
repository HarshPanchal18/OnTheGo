package main

import (
	"fmt"
	"sync"
	"time"
)

func mutex() {
	mutex := new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func(i int) {
			mutex.Lock()

			fmt.Println(i, " start")
			time.Sleep(time.Second)
			fmt.Println(i, " stop")

			mutex.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}