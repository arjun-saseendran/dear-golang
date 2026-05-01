package main

import "fmt"

func main() {

	process()
	fmt.Println("Return from process")

}

func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered", r)
		}
	}()
	fmt.Println("Start process")
	panic("Something went wrong!")
}
