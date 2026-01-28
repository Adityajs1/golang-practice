package main

import "fmt"

func tempo(celsius float64) {
    fahrenheit := (celsius * 9 / 5) + 32
    fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
}