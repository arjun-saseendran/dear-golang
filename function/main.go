package main

import (
	"fmt"
)

func main() {
	sumResult := sum(5, 8)
	mulResult := multiple(5, 8.9)
	divResult := divide(5, 3)

	fmt.Println(sumResult)
	fmt.Println(mulResult)
	fmt.Println(divResult)
	greet(display, "maria")

}

func sum(num1, num2 int) int {
	return num1 + num2
}

func greet(view func(string), data string) {
	view(data)
}

func display(name string) {
	fmt.Println("hello,", name)
}

var multiple = func(num1, num2 float64) float64 {
	return num1 * num2

}

func divide(num1, num2 float64) float64 {
	return num1 / num2
}
