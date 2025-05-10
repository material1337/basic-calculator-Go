package main

import "fmt"

func main() {
	var first, second float64
	var oper string
	fmt.Printf("Enter first number:")
	fmt.Scan(&first)
	fmt.Printf("Enter second number:")
	fmt.Scan(&second)
	fmt.Printf("Enter operation type:")
	fmt.Scan(&oper)

	switch oper {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "/":
		if second == 0 {
			fmt.Println("error:divide on null")
		} else {
			fmt.Println(first / second)
		}
	default:
		fmt.Println("error:unknown operation type")
	case "*":
		fmt.Println(first * second)
	}
}
