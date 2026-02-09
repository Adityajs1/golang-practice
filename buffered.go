package main

import "fmt"


func buffered(){
	ch := make(chan string)

	go func(){
		ch <- "data"
	}()

	val := <- ch;
	fmt.Println(val)
}