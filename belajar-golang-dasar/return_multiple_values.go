package main

import "fmt"

func getFullName() (string, string) {
    return "Ade", "Aulian"
}

func main() {
    // firstName, lastName := getFullName()
    // fmt.Println(firstName, lastName)


    // contoh kode program menghiraukan return value
    firstName, _ := getFullName()
    fmt.Println(firstName)
}