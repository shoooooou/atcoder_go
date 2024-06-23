package main

import (
	"fmt"
	"time"
)

func goroutine(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		fmt.Printf("i:%v,v:%v\n",i,v)
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	ss := []int{2, 4, 6, 8, 10}
	c := make(chan int)

	go goroutine(s, c)
	go goroutine(ss, c)
	x := <-c
	y := <-c
	fmt.Println(x)
	fmt.Println(y)
}
