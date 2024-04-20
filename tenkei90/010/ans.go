package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	var sum1 int
	var sum2 int
	for i := 0; i < q; i++ {
		scanner.Scan()
		start, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		end, _ := strconv.Atoi(scanner.Text())
		sum1 = score1[end] - score1[start-1]
		sum2 = score2[end] - score2[start-1]
		fmt.Printf("%d %d\n", sum1, sum2)
	}

}
