package main

import "fmt"

func main() {
	multiArray := [5][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println("The mutil dimensional array is:", multiArray)
	nums := [5]int{1, 4, 2, 5, 3}
	newNum := &nums
	fmt.Println("New array is:", newNum)
	pointeValue := *newNum
	fmt.Println("The value of pointer *newSum", pointeValue)
}
