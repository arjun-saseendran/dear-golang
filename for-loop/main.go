package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	for index, num := range nums {
		fmt.Printf("Index is: %v\n Value is: %v\n", index, num)
	}
	for i := 0; i < 5; i++ {
		fmt.Println("No is:", i)
	}

	for i := range 10 {
		fmt.Println("Range is:", i)
	}
k := 1

for k < 10 {
	fmt.Print(k)
	k++
}
}
