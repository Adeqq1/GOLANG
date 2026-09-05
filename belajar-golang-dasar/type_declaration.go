package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAde NoKTP = "1111111111111"

	var contoh string = "22222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpAde)
	fmt.Println(contohKtp)
}