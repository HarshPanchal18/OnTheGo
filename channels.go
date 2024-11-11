package main

import (
	"fmt"
	"time"

)

// Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
// A channel type is represented with the keyword `chan` followed by the type of the things that are passed on the channel.
// The <- (left arrow) operator is used to send and receive messages on the channel.

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <- c // receive a message and store it in `msg`
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/*
When pinger attempts to send a message on the channel it will wait until printer is ready to receive the message.
This is known as blocking.
*/