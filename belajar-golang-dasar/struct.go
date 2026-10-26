package main

import "fmt"

type Customer struct {
    Name, Address string
    Age           int
}

func (customer Customer) sayHello(name string) {
    fmt.Println("Hello", name, "my name is", customer.Name)
}


func main() {
    var ade Customer
    fmt.Println(ade)

    ade.Name = "Ade Rifqy Aulian"
    ade.Address = "Talang Pantai"
    ade.Age = 30
    fmt.Println(ade)
    fmt.Println(ade.Name)
    fmt.Println(ade.Address)
    fmt.Println(ade.Age)

    joko := Customer {
        Name: "Joko",
        Address: "Indonesia",
        Age: 30,
    }
    fmt.Println(joko)

    budi := Customer{"Budi", "Indonesia", 30}
    fmt.Println(budi)

    budi.sayHello("Agus")
    ade.sayHello("Agus")
    joko.sayHello("Agus")
}