package main

import ("fmt")


func selec() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        ch1 <- "from ch1"
    }()
    
    go func() {
        ch2 <- "from ch2"
    }()
    
    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
    // whichever sends first will be received
}