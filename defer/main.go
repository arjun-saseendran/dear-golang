package main

import "fmt"

func main() {
	defer fmt.Println("Defer result 1")
	defer fmt.Println("Defer result 2")
	defer fmt.Println("Defer result 3")
	fmt.Println("End of the main function")
}
