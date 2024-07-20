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

func sum(a int) int {
	return a * (a + 1) / 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf:=make([]byte,0, 60*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s := scanner.Text()
	oFlag := false
	xFlag := false
	subNum := 0
	cnt := 0
	for _, char := range s {
		if char == 'o' {
			oFlag = true
			if xFlag {
				xFlag = false
				subNum += sum(cnt)
				cnt = 0
			}
			cnt++
		} else {
			xFlag = true
			if oFlag {
				oFlag = false
				subNum += sum(cnt)
				cnt = 0
			}
			cnt++
		}
	}
	if cnt != 0 {
		subNum += sum(cnt)
	}
	// fmt.Println(sum(n),subNum)
	fmt.Println(sum(n) - subNum)
}
