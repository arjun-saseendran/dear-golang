package main

import "fmt"

func main() {
	num := 5

	if num%2 == 0 {
		fmt.Println("It is an even number.")
	} else if num%2 != 0 {
		fmt.Println("It is an odd number.")
	} else {
		fmt.Println("Please enter a valid number!")
	}
}
