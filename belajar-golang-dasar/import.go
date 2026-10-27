package main

import (
	"fmt"
	"belajar-golang-dasar/helper"
)

func main() {
	result := helper.SayHello("Ade")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.version) // tidak bisa diakses
	fmt.Println(helper.sayGoodBye("Ade")) // tidak bisa diakses
}