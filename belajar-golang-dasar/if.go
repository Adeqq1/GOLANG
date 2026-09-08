package main

import "fmt"

func main(){
	name := "Ade"

	if name == "Ade" {
		fmt.Println("Hello Ade")
	}else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}