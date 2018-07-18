package main

import (
    "fmt"

    "github.com/vevenhar/go-playground/io"
)

// the quintessential "Hello, World" program
func main() {
    fmt.Printf("Hello, World\n")

    fileName := "data/data1.csv"
    io.ReadCSV(fileName)
}
