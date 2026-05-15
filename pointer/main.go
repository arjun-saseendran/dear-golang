package main

import "fmt"

func main(){
	var ptr *int
	var num int = 10
	ptr = &num

	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&num)
	
}