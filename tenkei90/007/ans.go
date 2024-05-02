package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	sort.Ints(a)

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	var b int
	// time.Sleep(2 * time.Second)  // 3秒間スリープ

	for i := 0; i < q; i++ {
		scanner.Scan()
		b, _ = strconv.Atoi(scanner.Text())
		index := sort.Search(n, func(i int) bool {
			return a[i] >= b
		})
		if index == n {
			fmt.Println(b - a[n-1])
			continue
		}
		if index == 0 {
			fmt.Println(a[0] - b)
			continue
		}
		if abs(a[index]-b) < abs(a[index-1]-b) {
			fmt.Println(abs(a[index] - b))
		} else {
			fmt.Println(abs(a[index-1] - b))
		}
	}
}
