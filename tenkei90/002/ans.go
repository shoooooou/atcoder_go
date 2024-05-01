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
func formatBitStr(num int64, bitLength int) string {
	binaryStr := strconv.FormatInt(num, 2)

	return strings.Repeat("0", bitLength-len(binaryStr)) + binaryStr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	// nが奇数の場合は即時リターン
	if n%2 == 1 {
		return
	}

	binaryStr := strconv.FormatInt(int64((1<<n)-1), 2)
	binaryInt, _ := strconv.ParseInt(binaryStr, 2, 64)
	cntLeft := 0
	cntRight := 0
	displayStr := ""
	for i := binaryInt; i > 0; i-- {
		cntLeft = 0
		cntRight = 0
		displayStr = ""
		binaryStr = formatBitStr(i, n)
		for _, s := range binaryStr {
			if s == '1' {
				displayStr += "("
				cntLeft++
			} else {
				displayStr += ")"
				cntRight++
			}
			if cntLeft < cntRight {
				break
			}
		}
		if cntLeft == cntRight {
			fmt.Println(displayStr)
		}
	}
}
