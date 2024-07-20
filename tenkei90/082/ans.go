package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// スライスの合計値を算出する
func sumSlice(numbers []int64) int64 {
	var sum int64 = 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// 最大公約数を算出する
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 絶対値を算出する
func abs(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

// スライスの要素を入れ替える
func swap(array []int64, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// 順列を生成する
func generate(n int, array []int64, result *[][]int64) {
	if n == 1 {
		tmp := make([]int64, len(array))
		copy(tmp, array)
		*result = append(*result, tmp)
	} else {
		for i := 0; i < n; i++ {
			generate(n-1, array, result)
			if n%2 == 1 {
				swap(array, 0, n-1)
			} else {
				swap(array, i, n-1)
			}
		}
	}
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

// 総和を算出する
func sum(a int64) int64 {
	return (a % mod) * ((a + 1) % mod) / 2 % mod
}

const mod = 1000000007

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	l, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	r, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	var ans int64 = 0
	for i := int64(len(strconv.FormatInt(l, 10))); i <= int64(len(strconv.FormatInt(r, 10))); i++ {
		ans += ((sum(min(r, int64(math.Pow10(int(i)))-1)) - sum(max(l-1, int64(math.Pow10(int(i-1)))-1))) % mod) * i % mod
	}

	fmt.Println(ans)
}
