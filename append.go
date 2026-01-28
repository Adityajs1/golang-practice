package main

import "fmt"

func appended(){
  integers := []int{};
  integers = append(integers, 10, 20, 30);
  fmt.Println(integers)
}