package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	slice := nums[1:3]
	fmt.Println("Slice data is:", slice)
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("Copied slice data:", sliceCopy)
	isEqual := slices.Equal(slice, sliceCopy)
	fmt.Println("Equal check result:", isEqual)
	var nilSlice []int
	fmt.Println("Nil slice data:", nilSlice)
	slice = append(slice, 7, 8, 9)
	fmt.Println("Updated slice data:", slice)
	slice[0] = 7
	fmt.Println("Updated slice data:", slice)
	for index, value := range slice {
		fmt.Printf("Index is: %d Value is:%v\n", index, value)
	}

	newSlice := slice
	newSlice[0] = 5
	fmt.Println("Updated new slice data is:", newSlice)
	fmt.Println("Old slice data is:", slice)
	isEqual = slices.Equal(slice, sliceCopy)
	fmt.Println("Equal check result:", isEqual)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two dimentianal slice data:", twoD)
	fmt.Println("The length of the array is:", len(slice))
	fmt.Println("The capacity of the slice is:", cap(slice))
}
