package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func person(){
  s := Person{
	Name: "Alice",
	Age: 25,
  }
  s1 := Person{
	Name: "Bob",
	Age: 30,
  }

 fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
fmt.Printf("Name: %s, Age: %d\n", s1.Name, s1.Age)

}