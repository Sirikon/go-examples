package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//exampleA()
	//exampleB()
	exampleC()
}

func exampleA() {
	messages := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "One"
		time.Sleep(1 * time.Second)
		messages <- "Two"
		time.Sleep(1 * time.Second)
		messages <- "Three"
	}()

	fmt.Println("Waiting for ONE")
	fmt.Println(<-messages)
	fmt.Println("Waiting for TWO")
	fmt.Println(<-messages)
	fmt.Println("Waiting for THREE")
	fmt.Println(<-messages)
}

func exampleB() {
	// Change the buffer length and see what happens
	messages := make(chan string, 0)

	go func() {
		messages <- "One"
		fmt.Println("Inserted One")
		messages <- "Two"
		fmt.Println("Inserted Two")
		messages <- "Three"
		fmt.Println("Inserted Three")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(<-messages)
	time.Sleep(1 * time.Second)
	fmt.Println(<-messages)
	time.Sleep(1 * time.Second)
	fmt.Println(<-messages)
}

func backgroundProcess(end chan bool) {
	fmt.Println("Working on the back")
	time.Sleep(1 * time.Second)
	fmt.Println("Working on the back again")

	end <- true
}

func exampleC() {
	end := make(chan bool, 1)

	go backgroundProcess(end)

	fmt.Println("Working on the front")
	fmt.Println("Waiting for background")

	result := <-end
	fmt.Println("Result: " + strconv.FormatBool(result))
}
