package main

import "fmt"

func endApp() {
    fmt.Println("END APP")
    message := recover()
    fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
    defer endApp()
    if error {
        panic("Ups Error")
    }
}



func main() {
    runApp(true)
    fmt.Println("Ade Rifqy Aulian")
}