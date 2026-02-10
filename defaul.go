package main

import (
	"fmt"
)

// here there is no data availaible that's why it is always printing "no data availaible"

func defaul() {
    ch := make(chan int)
    
    select {
    case val := <-ch:
        fmt.Println("received:", val)
    default:
        fmt.Println("no data available") // executes immediately
    }
}