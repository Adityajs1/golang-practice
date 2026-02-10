package main

import ("fmt")

func goaxiom(){
	var ch chan int
	ch = make(chan int)
    close(ch)

	ch = make(chan int)
    close(ch)
    val := <-ch
    fmt.Println(val)
}