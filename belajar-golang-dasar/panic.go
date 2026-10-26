package main

import "fmt"

func endApp() {
    fmt.Println("End APP")
}

func runApp(error bool) {
    defer endApp()

    if error {
        panic("Ups Error")
    }
}

func main() {
    runApp(true)
}