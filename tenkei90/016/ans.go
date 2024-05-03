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

// 絶対値を算出する
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	ans := 10000
	// a円、b円、c円の最小の組み合わせでn円を作る
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000-i; j++ {
			// n円以上の場合はスキップ
			if n < i*b+j*a {
				break
			}
			if (n-i*b-j*a)%c == 0 {
				// fmt.Println("a", j, "b", i, "c", (n-(i*b)-(j*a))/c)
				ans = min(ans, i+j+(n-i*b-j*a)/c)
			}
		}
	}
	fmt.Println(ans)
}
