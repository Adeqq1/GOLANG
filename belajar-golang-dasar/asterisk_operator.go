package main

import "fmt"

type Address struct {
    City, Province, Country string
}

func main() {
    var address1 Address = Address{"Bungo", "Jambi", "Indonesia",}
    var address2 *Address = &address1  

    address2.City = "Tebo"
    fmt.Println(address1) // ikut berubah 
    fmt.Println(address2) // berubah

    // address2 = &Address {"Padang", "Sumatera Barat", "Indonesia"}
    *address2 = Address {"Padang", "Sumatera Barat", "Indonesia"}
    fmt.Println(address1)
    fmt.Println(address2)
}