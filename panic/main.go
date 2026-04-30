package main

import "fmt"

func main() {
	defer fmt.Println("Defer result 1")
	defer fmt.Println("Defer result 2")
	process(-4)

}

func process(num int) {
	if num < 0 {
		panic("Number must be greaterthan zero!")
	}
	defer fmt.Println("End of the function!")
}
