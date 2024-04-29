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
	n, m, a, b := dataInput()
	cnt := make([]int, n+1)
	for i := 0; i < m; i++ {
		if a[i] > b[i] {
			cnt[a[i]]++
		} else {
			cnt[b[i]]++
		}
	}
	ans := 0
	for _, cntAns := range cnt {
		if cntAns == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
func dataInput() (int, int, []int, []int) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	return n, m, a, b
}
