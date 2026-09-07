package main

import "fmt"

func main(){
	// var person map[string]string = map[string]string{}
	// person["name"] = "Ade"
	// person["address"] = "Talang Banjar"

	person := map[string]string{
		"name": "Ade",
		"address": "Danau Laman",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Ade"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}