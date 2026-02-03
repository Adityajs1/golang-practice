package main

import "fmt"

func ageMap(){
	var ages = map[string]int {
		"Aditya" : 19,
		"Yash" : 20,
		"Dheeraj" : 19,
	}
	for name, age := range ages{
		fmt.Printf("%s is %d years old\n", name, age)
	}
}