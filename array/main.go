package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "Pratham"
	name[1] = "Praveen"
	fmt.Println(name)
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	var price [5]int
	fmt.Println(price)
}
