package main

import "fmt"

func init(){
	fmt.Println("Message from init function")
}

func init(){
	fmt.Println("Message from second init function")
}

func main() {
	fmt.Println("Message from main function")
}
