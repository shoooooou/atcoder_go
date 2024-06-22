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

// スライスの要素を入れ替える
func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// 順列を生成する
func generate(n int, array []int, result *[][]int) {
	if n == 1 {
		tmp := make([]int, len(array))
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

func main() {
	const mod = 1000000007
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	a := make([]int, 6)
	ans := 1
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < 6; j++ {
			scanner.Scan()
			a[j], _ = strconv.Atoi(scanner.Text())
			sum += a[j] % mod
		}
		ans = (ans * sum) % mod
	}
	fmt.Println(ans)
}
