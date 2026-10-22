package main

import "fmt"

func getHello(name string) string {
    hello := "Hello " + name
    return hello
}

func main() {
    result := getHello("Ade")
    fmt.Println(result)
    
    fmt.Println(getHello("Kink"))
    fmt.Println(getHello("Aulian"))
}