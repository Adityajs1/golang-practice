package main

import "fmt"

func channelstwo(){
	makeChannel := make (chan string)
	anotherChannel := make (chan string)

	go func(){
		makeChannel <- "data"
	}()

    go func(){
		anotherChannel <- "cow"
	}()
	
	//SELECT statement
	select{
	case msgFromMyChannel := <-makeChannel:
	fmt.Println(msgFromMyChannel)
	case msgFromMyChannel := <-anotherChannel:
	fmt.Println(msgFromMyChannel)
	}
}