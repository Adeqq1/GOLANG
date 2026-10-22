package main

import "fmt"

func sayHelloTo (firstName string, middleName string, lastName string, fullName string) {
    fmt.Println ("Hello", firstName, middleName, lastName, fullName)
}

func main() {
    sayHelloTo("Ade", "Rifqy", "Aulian", "Ade Rifqy Aulian")
    sayHelloTo("Anak", "Baik", "Aulian", "Adeqq1")
}