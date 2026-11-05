package main

import "fmt"

func main() {
    firstName := "Ade"
    lastName := "Aulian"

    fmt.Println("Hello '", firstName, lastName, "'")

    fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}