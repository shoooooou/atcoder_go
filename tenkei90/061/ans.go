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
	q, _ := strconv.Atoi(scanner.Text())
	query := make([][2]int, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		query[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		query[i][1], _ = strconv.Atoi(scanner.Text())
	}
	slice := make([]int, 500000*2+1)
	st:=500000
	ls:=500000
	for _, q := range query {
		if q[0] == 1 {
			slice[st]=q[1]
			st--
		} else if q[0] == 2 {
			slice[ls+1]=q[1]
			ls++
		} else {
			fmt.Println(slice[st+q[1]])
		}
	}

}
