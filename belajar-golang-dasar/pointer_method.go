package main

import "fmt"

type Man struct {
    Name string
}

func (man *Man) Married() {
    man.Name = "Mr. " + man.Name
}

func main() {
    ade := Man{"Ade"}
    ade.Married()

    fmt.Println(ade.Name)
}