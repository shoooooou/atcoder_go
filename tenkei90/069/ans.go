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

const mod = 1e9 + 7

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	if n == 1 {
		fmt.Println(k)
		return
	}
	if n == 2 {
		if k == 1 {
			fmt.Println(0)
		} else {
			fmt.Println((k * (k - 1)) % mod)
		}
		return
	}
	if k <= 2 {
		fmt.Println("0")
		return
	}

	var ans int
	ans = ((k * (k - 1) % mod) * powMod(k-2, n-2)) % mod
	fmt.Println(ans)
}
func powMod(x int, y int) int {
	v, p := 1, x
	for ; y > 0; y >>= 1 {
		if y&1 == 1 {
			v = (v * p) % mod
		}
		p = (p * p) % mod
	}
	return v
}
