package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// スライスの合計値を算出する
func sumSlice(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// 最大公約数を算出する
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	if h == 1 || w == 1 {
		fmt.Println(h * w)
		return
	}
	fmt.Println((h+1)/2 * (w+1)/2)
}
