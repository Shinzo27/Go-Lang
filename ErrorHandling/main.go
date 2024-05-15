package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	divide, err := divide(10, 4)
	fmt.Println(divide)
	if err != nil {
		fmt.Println(err)
	}
}
