package main

import "fmt"

func main() {
	// name := "Maria"
	num := 5

	checkType(num)
	color := "Blue"

	switch color {
	case "Blue":
		fmt.Println("It is a color")
		fallthrough
	case "Yellow":
		fmt.Println("It is a color")
	case "Red":
		fmt.Println("It is a color")
	default:
		fmt.Println("It is a fallthrough condition")
	}
}

func checkType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("It is an integer type")
	case float32:
		fmt.Println("It is a float type")
	case bool:
		fmt.Println("It is a boolean type")
	default:
		fmt.Println("Invalid type!")
	}
}
