package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	age := 25
	var name string
	height := 6.100

	fmt.Println("Age:", age, "Name:", name, "height:", height)
	fmt.Printf("Height is %.2f \n", height)
	fmt.Printf("Height is %d \n", age)
	// fmt.Print("Enter Your name: ")
	// fmt.Scan(&name)
	// fmt.Println("Hello Mr.", name)

	reader := bufio.NewReader(os.Stdin)
	firstName, _ := reader.ReadString('\n')
	fmt.Println("Hello Mr. ", firstName)
}
