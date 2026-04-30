package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	totalSum := variadic(nums...)
	fmt.Println("The total sum is:", totalSum)
}

func variadic(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
