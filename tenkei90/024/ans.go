package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 最大公約数を算出する
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// スライスの合計値を算出する
func sumSlice(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	diffSum := 0
	for i := 0; i < n; i++ {
		diffSum += abs(a[i] - b[i])
	}

	if k < diffSum {
		fmt.Println("No")
		return
	}
	if abs(k-diffSum)%2 != 0 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
