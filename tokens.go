package main

import (
	"fmt"
)


func tokens(){
	done := make(chan bool)

	go func(){
		fmt.Println("I am Working...")
		done <- true
	}()
	<-done 
    fmt.Println("Work complete!")


}