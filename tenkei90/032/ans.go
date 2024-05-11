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

// swap values in array
func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// generate permutations using Heap's Algorithm
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
	section := make([][]int, n)
	for i := range section {
		section[i] = make([]int, n)
	}

	for i := range section {
		for j := range section[i] {
			scanner.Scan()
			section[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	notFriends := make([][]bool, n)
	for i := 0; i < n; i++ {
		notFriends[i] = make([]bool, n)
	}

	for i := 0; i < q; i++ {
		// fmt.Println("asdf;lhjadskl;fja")
		scanner.Scan()
		num1, _ := strconv.Atoi(scanner.Text())
		num1--
		scanner.Scan()
		num2, _ := strconv.Atoi(scanner.Text())
		num2--
		notFriends[num1][num2] = true
		notFriends[num2][num1] = true

	}

	permutationRow := make([]int, n)
	for i := range permutationRow {
		permutationRow[i] = i
	}
	var permutation [][]int
	generate(len(permutationRow), permutationRow, &permutation)

	ans := -1
	for permuVals := range permutation {

		if q != 0 && !checkNotFriends(permutation[permuVals], notFriends) {
			continue
		}
		// ans = -1
		sum := 0
		for i, permuVal := range permutation[permuVals] {
			sum += section[i][permuVal]
		}
		if ans == -1 || sum < ans {
			ans = sum
		}

	}
	fmt.Println(ans)
}

// 仲違いの組み合わせがないかを判定する
func checkNotFriends(permutation []int, notFriends [][]bool) bool {
	for i := 0; i < len(permutation)-1; i++ {
		if notFriends[permutation[i]][permutation[i+1]] {
			return false
		}
	}
	return true
}
