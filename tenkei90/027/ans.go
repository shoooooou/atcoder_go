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
	n, _ := strconv.Atoi(scanner.Text())
	strDup := make(map[string]bool)
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		_, exists := strDup[str]
		if !exists {
			strDup[str] = true
			fmt.Println(i + 1)
		}
	}
}
