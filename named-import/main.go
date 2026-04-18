package main

import (
	f "fmt"
	h "net/http"
)


func main(){
	res, err := h.Get("https://dummyjson.com/users")
	if err != nil {
		f.Println("error fetching users.")
		return
	}
	defer res.Body.Close()
	f.Println("data status is: ", res.Status)
}