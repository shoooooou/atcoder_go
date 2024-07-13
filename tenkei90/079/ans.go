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
	h, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())

	a := make([][]int, h)
	b := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		b[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			scanner.Scan()
			a[i][j], _ = strconv.Atoi(scanner.Text())
			// fmt.Printf("%d ",a[i][j])
		}
	}
	// fmt.Println("")
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			scanner.Scan()
			b[i][j], _ = strconv.Atoi(scanner.Text())
			// fmt.Printf("%d ",b[i][j])
		}
	}
	ans := 0
	var subNum int
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			if a[i][j] == b[i][j] {
				continue
			} else if a[i][j] > b[i][j] {
				subNum = a[i][j] - b[i][j]
				a[i][j] -= subNum
				a[i+1][j] -= subNum
				a[i][j+1] -= subNum
				a[i+1][j+1] -= subNum
			} else {
				subNum = b[i][j] - a[i][j]
				a[i][j] += subNum
				a[i+1][j] += subNum
				a[i][j+1] += subNum
				a[i+1][j+1] += subNum
			}
			ans += subNum
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] != b[i][j] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
	fmt.Println(ans)

}
