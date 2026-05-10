package main

import "os"


func main(){
	println("Program started!")
	os.Exit(1)
	defer println("Hello from defer will not executed")
}