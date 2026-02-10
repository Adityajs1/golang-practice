package main

import (
	"fmt"
)

func valok(){
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	close(ch)

	val, ok := <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)

}
