package main

import(
	"fmt"
)
import (
	"time"
)


func standardtime(){
	ticker := time.Tick(1 * time.Second)
    for i := 0; i < 3; i++ {
        <-ticker // blocks for 1 second each time
        fmt.Println("tick", i)
    }
    
    // time.After - sends once after duration
    <-time.After(2 * time.Second)
    fmt.Println("2 seconds passed")
    
    // time.Sleep - just blocks
    time.Sleep(1 * time.Second)
    fmt.Println("slept for 1 second")
}