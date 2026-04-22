package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := 5
	num2 := 9
	var result int

	result = num1 + num2
	fmt.Println("Addition:", result)
	result = num1 - num2
	fmt.Println("Subtraction:", result)
	result = num1 * num2
	fmt.Println("Multiplication:", result)
	result = num1 / num2
	fmt.Println("Devisiton:", result)
	result = num1 % num2
	fmt.Println("Reminder is:", result)

	const num3 float64 = 10 / 3.0

	fmt.Println("Divided result:", num3)

	// Overflow with signed integers.
	var maxInt int64 = 9223372036854775807
	fmt.Println("Max int64 value is:", maxInt)

	maxInt = maxInt + 1
	fmt.Println("Overflow after adding 1 to maxInt:", maxInt)
	
	// Overflow with unsigned integers.
	var uMaxint uint64 = 18446744073709551615
	fmt.Println("Max uint64 value is:", uMaxint)
	uMaxint = uMaxint + 1
	fmt.Println("Overflow after adding 1 to uMaxInt",uMaxint)
	
	// Underflow with float.
	var smallFloat float64 = 1.0e-323
	fmt.Println("Float64 for small value is:", smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Underflow value after dividing with max float64 value is:", smallFloat)

}
