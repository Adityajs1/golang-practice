package main

import ("fmt")


func send(ch chan<- int) { // can only send
    ch <- 42
}

func receive(ch <-chan int) { // can only receive
    val := <-ch
    fmt.Println(val)
}

func only() {
    ch := make(chan int)
    go send(ch)
    receive(ch)
}