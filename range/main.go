package main

import "fmt"

func main() {

	name := "Arjun"

	for index, value := range name {
		fmt.Println("The index is: ", index, "\nThe value is:", value)
		fmt.Printf("The index is: %v\nvalue is: %c\n",index,value)
	}

}
