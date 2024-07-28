package main

import "fmt"

func main() {
	fmt.Println("Learn Go Lang!")
	var name string = "Pratham"
	var version = 1.2
	fmt.Println(name)
	fmt.Println(version)

	var money int = 20
	fmt.Println(money)
	var ruppees float64 = 10.20
	fmt.Println(ruppees)

	const pi = 3.14
	fmt.Println(pi)

	//declare variable without type
	person := 123
	fmt.Println(person)

	//package can be uploaded if it is capitalized
	var Public = "Public"
	//package can be uploaded if it is capitalized
	var private = "private"

	fmt.Println(Public)
	fmt.Println(private)

	var message = "Hello world! we are back with go lang!"
	fmt.Println(message)
}
