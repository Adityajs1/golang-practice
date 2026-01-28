package main

import "fmt"

func returns(a, b int)int{
  return a + b
}

func re(){
	result := returns(2, 3)
	fmt.Println("Sum:",result)
}