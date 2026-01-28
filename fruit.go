package main

import "fmt"

func fruit(){
	fruits := []string{"Apple", "Mango","Banana", "Grapes", "Orange"}
	for _, fruit := range fruits{
     fmt.Println(fruit)
	}
	}
