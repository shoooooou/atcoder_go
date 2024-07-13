package main

import (
	"bufio"
	"fmt"
	"math"
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
	// 3以下の場合は即時リターン
	if n <= 3 {
		fmt.Println("0")
		return
	}
	// ルートnまでの素数をスライスで算出する
	primes := generatePrimes(n)
	fmt.Println("primes", primes)
	// // nが素数の場合は即時リターン
	// for i := 1; i < len(primes); i++ {
	// 	if n%primes[i] == 0 && i == len(primes)-1 {
	// 		fmt.Println("0")
	// 		return
	// 	}
	// }

	cnt := 0
	find := true
	for find {
		for i, v := range primes {
			if n%v == 0 {
				cnt++
				n /= v
				continue
			}
			if i == len(primes)-1 {
				find = false
			}
		}
	}

	fmt.Println(n - 1)
}

func generatePrimes(n int) []int {
	sqrtNum := int(math.Sqrt(float64(n)))
	boolSlice := make([]bool, sqrtNum+1)
	for i := 2; i < sqrtNum+1; i++ {
		boolSlice[i] = true
	}
	p := 2
	for p*p <= n {
		if boolSlice[p] {
			for i := p * 2; i*i <= n; i += p {
				boolSlice[i] = false
			}
		}
		p++
	}
	var primeSlice []int
	for i := 2; i < len(boolSlice); i++ {
		if boolSlice[i] {
			primeSlice = append(primeSlice, i)
		}
	}
	return primeSlice
}
