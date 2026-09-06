package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "Ade"
	names[1] = "Rifqy"
	names[2] = "Aulian"

	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[0])

	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}