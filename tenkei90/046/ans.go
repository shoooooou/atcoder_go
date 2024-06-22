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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var chkNum int
	a := make([]int, 46)
	for i := 0; i < n; i++ {
		scanner.Scan()
		chkNum, _ = strconv.Atoi(scanner.Text())
		a[chkNum%46]++
	}
	b := make([]int, 46)
	for i := 0; i < n; i++ {
		scanner.Scan()
		chkNum, _ = strconv.Atoi(scanner.Text())
		b[chkNum%46]++
	}
	c := make([]int, 46)
	for i := 0; i < n; i++ {
		scanner.Scan()
		chkNum, _ = strconv.Atoi(scanner.Text())
		c[chkNum%46]++
	}

	ans :=0
	for ai,av :=range a{
		for bi,bv:=range b{
			for ci,cv:=range c {
				if (ai+bi+ci)%46==0{
					ans+=av*bv*cv
				}
			}
		}
	}

	fmt.Println(ans)
}
