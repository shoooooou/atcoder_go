package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const mod = 1000000007

// 総和を算出する
func sum(a int64) int64 {
	return ((a % mod) * ((a + 1) % mod) / 2) % mod
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	l, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	r, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	var ans int64 = 0
	for i := int64(len(strconv.FormatInt(l, 10))); i <= int64(len(strconv.FormatInt(r, 10))); i++ {
		low := max(l, int64(math.Pow10(int(i-1))))
		high := min(r, int64(math.Pow10(int(i)))-1)
		if low <= high {
			ans += ((sum(high) - sum(low-1) + mod) % mod) * i % mod
			ans %= mod
		}
	}

	fmt.Println(ans)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
