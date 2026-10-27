package main

import "fmt"

type Address struct {
    City, Province, Country string
}

func main() {
    // address1 := Address{
    //     "Bungo", "Jambi", "Indonesia",
    // }
    // address2 := address1 // copy value address1

    // address2.City = "Tebo"
    // fmt.Println(address1) // tidak berubah
    // fmt.Println(address2)

    var address1 Address = Address{
        "Bungo", "Jambi", "Indonesia",
    }
    var address2 *Address = &address1  

    address2.City = "Tebo"
    fmt.Println(address1) // ikut berubah 
    fmt.Println(address2) // berubah
}