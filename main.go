package main

import (
	"fmt"
	"project/anggunpermata/calculator"
)

func main() {
	fmt.Println(calculator.Addition(3, 3))
	fmt.Println(calculator.Substraction(3, -1))
	fmt.Println(calculator.Multiplication(4, -2))
	fmt.Println(calculator.Division(2, -1))
}
