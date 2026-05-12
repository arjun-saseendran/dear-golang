package main

import "fmt"

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	reset := counter()
	fmt.Println(reset())

	decrement := func() func(int) int {
		count := 10
		return func(num int) int {
			count -= num
			return count
		}
	}()

	fmt.Println(decrement(1))
	fmt.Println(decrement(2))
	fmt.Println(decrement(3))

}

func counter() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}
