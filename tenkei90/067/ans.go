package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	nStr := scanner.Text()
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	var nInt int64
	for i := 0; i < k; i++ {
		// 10進数の数値に変換
		nInt, _ = strconv.ParseInt(nStr, 8, 64)
		// 9進数の文字列に変換
		nStr = strconv.FormatInt(nInt, 9)
		// 8を5に置換
		nStr = strings.ReplaceAll(nStr, "8", "5")
	}
	fmt.Println(nStr)

}
