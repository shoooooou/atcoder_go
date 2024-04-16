package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 公約数を算出する
// TODO: スニペット化する
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
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	gcdValue := gcd(gcd(a, b), c)

	ans := (a/gcdValue - 1) + (b/gcdValue - 1) + (c/gcdValue - 1)
	fmt.Println(ans)
}
