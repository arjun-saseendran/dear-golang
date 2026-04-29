package main

import "fmt"

func main() {
	sum, mul := calcu(5, 8)
	fmt.Println("Sum is: ", sum)
	fmt.Println("Multiple is: ", mul)
}

func calcu(num1, num2 int) (sum, mul int) {
	return num1 + num2, num1 * num2
}
