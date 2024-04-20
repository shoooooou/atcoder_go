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
	H, _ := strconv.Atoi(scanner.Text())
	fmt.Println(H)
}