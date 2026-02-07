package main

import "fmt"

func forselect(){
	myChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars{
		select{
		case myChannel <- s:
		}
	}
	close(myChannel)

	for result := range myChannel{
      fmt.Println(result)
	}

}