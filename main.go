package main

import (
	"fmt"
	"strings"
)

//Goroutines
//keyword => go

func main2() {
	go count(7, 200, 50)
	go count(1, 20, 5)
	go count(10, 400, 100)
	start := 0
	stop := 50
	step := 2

	go func() {
		go count(start, stop, step)
	}()

	fmt.Scanln()

	//
}

func count(start, stop, step int) {
	fmt.Println("Starting at ", start)
	for i := start; i <= stop; i += step {
		fmt.Println(i)
	}
}

// Channels
//keyword => chan

func main3() {
	//var intCh chan int
	//intCh <- 12 //Send Operation sending data(12) to the intCh
	//_ = <-intCh //intCh Receive Operation,

	//Channel zero state => nil

	//Channel Initialization [Buffered or Unbuffered]
	ch1 := make(chan int) // unbuffered integer channel

	//1. ch1 => receiver blocked (x := <- ch1 will block)
	//2. ch1 => sender sends data into empty channel, receiver is still blocked
	//3. ch1 => receiver proceeds(is open), and sender is blocked

	go func() { ch1 <- 12 }()
	fmt.Println(<-ch1)

	//Buffered Channel
	ch2 := make(chan int, 4)
	//[2][21][22][23]

	//1. When channel is empty receiver blocks
	//2. When sender starts sending, receiver proceeds even though channel is not full yet
	//3. When channel is full or at capacity, sender blocks, receiver continues

	ch2 <- 2
	ch2 <- 21
	ch2 <- 22
	ch2 <- 23

	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)

	//Unidirectional Channels
	//<-
	//var inCh chan<- int  //receive only channel
	//var outCh <-chan int //send only channel

	ch3 := make(chan int, 10)
	makeEvenNums(4, ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println("Length of Ch3: ")
	fmt.Println(len(ch3))
	fmt.Println(cap(ch3))
}

func makeEvenNums(count int, in chan<- int) {
	for i := 0; i < count; i++ {
		in <- 2 * i
	}
}

//Channel length and capacity

//Closing a Channel
func main() {
	//ch1 := make(chan int, 4)
	//ch1 <- 2
	//ch1 <- 4
	//close(ch1)
	data := []string{
		"The yellow fish swims slowly in the water",
		"The world cup is gradually coming to a close",
		"The white bird of the great landmark location",
	}

	histogram := make(map[string]int) //{"string": 0}

	done := make(chan struct{})

	//split and count words
	go func() {
		defer close(done) //run this after all the code in this function has executed
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
	}()

	<-done

	for key, value := range histogram {
		fmt.Printf("%s\t (%d)\n", key, value)
	}

}
