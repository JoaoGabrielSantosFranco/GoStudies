package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

func main() {
	times4 := multiplier(2)
	fmt.Println(times4(4))
	fmt.Println(times4(3))
}
