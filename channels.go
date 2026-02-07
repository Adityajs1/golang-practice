package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func channels() {
    sayHello()
    
    // Goroutine - runs concurrently
    go sayHello()
    
    // Wait a bit so goroutine can finish
    time.Sleep(1 * time.Second)
    fmt.Println("Main function ends")
}

