package main

import (
    "fmt"
    "time"
)

func printMessage() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go printMessage() // Runs concurrently
    time.Sleep(time.Second) // Give time for goroutine to execute
}
