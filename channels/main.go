package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// basics()
	useChannel()
	useChannelLoop()
	bufferChannel()
	chickenRun()
}

// create, assign value, retrieve value, but deadlock
func basics() {
	fmt.Println("\n=== basics ===\n")
	var c = make(chan int)

	// assign value
	c <- 1

	// retrieve value
	var i = <-c

	fmt.Println(i)
}

func useChannel() {
	fmt.Println("\n=== use channel single value ===\n")
	var c = make(chan int)
	go process(c)
	fmt.Println(<-c)
}
func process(c chan int) {
	c <- 542
}

func useChannelLoop() {
	fmt.Println("\n=== use channel loop ===\n")
	var c = make(chan int)
	go processLoop(c)

	for i := range c {
		fmt.Println(i)
	}
}
func processLoop(c chan int) {
	// defer will run before func exit
	// close will let other processes use the channel
	defer close(c)

	for i := 0; i < 5; i++ {
		c <- i
	}
}

func bufferChannel() {
	fmt.Println("\n=== buffer channel ===\n")

	// buffer size 2: do 2 things at a time in a channel
	var c = make(chan int, 2)
	go processBuffer(c)

	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}
func processBuffer(c chan int) {
	defer close(c)

	for i := 0; i < 5; i++ {
		c <- i
	}

	// will exit in the middle of printing because work is completed in processBuffer
	fmt.Println("Exiting processBuffer")
}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func chickenRun() {
	fmt.Println("\n=== Chicken Run ===\n")

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}
func checkChickenPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}
func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		tofuPrice := rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// select case will run each block as a channel gives a value
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %s", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound a deal on tofu at %s\n", website)
	}
}
