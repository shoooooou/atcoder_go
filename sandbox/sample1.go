package main

import (
	"fmt"
)

func goroutine(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		fmt.Printf("i:%v,v:%v\n", i, v)
		sum += v
	}
	c <- sum
}

func main() {
	num := 10
	fmt.Printf("%b\n", num)
	num = num >> 1
	fmt.Printf("%b", num)
}
