package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) { // recieves the channel
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true 
	close(doneChan) //close the channel
}

// Goroutine & channels  - with this the program will executed based on which func compelets first 
func main() {
	done := make(chan bool) 
	go greet("How are you?", done)  // add go keyword to run the func concurrently 
	go slowGreet("How ... are ... you ...?", done)
	go greet("Nice to meet you!", done)
    

	for range done {  //for loop on a channel 
		// fmt.Println(doneChan)
	}
}
