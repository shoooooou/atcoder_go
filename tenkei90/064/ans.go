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
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	ans := calcHyoukoh(a)
	// fmt.Println(ans)

	for i := 0; i < q; i++ {
		scanner.Scan()
		l, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		l--
		r, _ := strconv.Atoi(scanner.Text())
		r--
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		if l != 0{
			fmt.Printf("left=bef:%v af:%v",abs(a[l-1]-a[l]),abs(a[l-1]-a[l]-v))
			ans+=(abs(a[l-1]-a[l]-v)-abs(a[l-1]-a[l]))
		}
		if r != (n-1){
			ans+=(abs(a[r]+v-a[r-1])-abs(a[r]-a[r-1]))
			fmt.Println("rans:",ans)
		}
		fmt.Println(ans)
	}

}
func calcHyoukoh(a []int) int {
	sum := 0
	for i := 0; i < len(a)-1; i++ {
		tmp := a[i] - a[i+1]
		if tmp < 0 {
			sum += -tmp
		} else {
			sum += tmp
		}
	}
	return sum
}
