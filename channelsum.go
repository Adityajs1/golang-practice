package main

import "fmt"

func calculator(a, b int, result chan int) {
    sum := a + b
    result <- sum  // Send answer back through channel
}

func channelsum() {
    resultChannel := make(chan int)
    
    go calculator(5, 3, resultChannel)  // Ask worker to calculate 5+3
    
    answer := <-resultChannel  // Wait for the answer
    fmt.Println("Answer:", answer)  // Answer: 8
}