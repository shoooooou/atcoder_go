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
	buf := make([]byte, 0, 60*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	score1 := make([]int, n+1)
	score2 := make([]int, n+1)
	for i := 0; i < n; i++ {
		scanner.Scan()
		classNum, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		score, _ := strconv.Atoi(scanner.Text())
		if classNum == 1 {
			score1[i+1] = score1[i] + score
			score2[i+1] = score2[i]
		} else {
			score2[i+1] = score2[i] + score
			score1[i+1] = score1[i]
		}
	}
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < q; i++ {
		var sum1 int
		var sum2 int
		scanner.Scan()
		start, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		end, _ := strconv.Atoi(scanner.Text())
		sum1 = score1[end] - score1[start-1]
		sum2 = score2[end] - score2[start-1]
		fmt.Printf("%d %d\n", sum1, sum2)
	}
}
