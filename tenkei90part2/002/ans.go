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
func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 60*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
func formatBitStr(num int64, bitLength int) string {
	// stringの2進数表記を取得
	binaryStr := strconv.FormatInt(num, 2)

	return strings.Repeat("0", bitLength-len(binaryStr)) + binaryStr
}

func main() {
	num := input()

	if num%2 == 1 {
		return
	}

	for bitKakko := int64((1 << num) - 1); bitKakko >= 0; bitKakko-- {
		cntLeft := 0
		cntRight := 0
		kakko := ""
		for _, bit := range formatBitStr(bitKakko, num) {
			if bit == '1' {
				kakko += "("
				cntLeft++
			} else {
				kakko += ")"
				cntRight++
			}
			if cntLeft < cntRight {
				break
			}
		}
		if cntLeft == cntRight {
			fmt.Println(kakko)
		}
	}
}
