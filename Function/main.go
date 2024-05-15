package main

import "fmt"

func simpleFunction() {
	fmt.Println("HEllo from function!")
}
func calculate(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Learning function")
	simpleFunction()
	fmt.Println(calculate(10, 20))
}
