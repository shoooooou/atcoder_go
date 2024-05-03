package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	// 生徒の入力
	students:=make([]int,n)
	for i:=0;i<n;i++{
		scanner.Scan()
		students[i], _ = strconv.Atoi(scanner.Text())
	}
	// 学校の入力
	schools:=make([]int,n)
	for i:=0;i<n;i++{
		scanner.Scan()
		schools[i], _ = strconv.Atoi(scanner.Text())
	}
	
	sort.Ints(students)
	sort.Ints(schools)

	var ans int
	for i:=0;i<n;i++{
		ans+=abs(students[i]-schools[i])
	}

	fmt.Println(ans)
}