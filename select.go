package main

import (
	"fmt"
	"time"

)

/*
`select` works like a `switch...case` but for channels.
This function prints “from 1” every 2 seconds and “from 2” every 3 seconds.
`select` picks the first channel that is ready and receives from it (or sends to it).
If more than one of the channels are ready then it randomly picks which one to receive from.
If none of the channels are ready, the statement blocks until one becomes available.
The select statement is often used to implement a timeout.
*/

func selections() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Channel 1"
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "Channel 2"
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)

			case msg2 := <- c2:
				fmt.Println(msg2)

			// `time.After`` creates a channel and after the given duration will send the current time on it.
			case <- time.After(time.Second):
				fmt.Println("timeout...")

			/*default: // happens immediately if none of the channels are ready.
				fmt.Println("Nothing ready!")*/

			}
		}
	}()

	// c3 := make(chan int, 1) // This creates a buffered channel with a capacity of 1.
	/*
	Normally channels are synchronous; both sides of the channel will wait until the other side is ready.
	A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.
	*/
}