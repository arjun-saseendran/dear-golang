package main

import "fmt"

func main() {
	var result = digitsOfSum(12)
	
	fmt.Println("Total sum is:", result)
}

func digitsOfSum(num int) int {
	if num < 10 {
		return num
	}
	return num%10 + digitsOfSum(num/10)
}
