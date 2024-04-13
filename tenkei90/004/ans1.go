package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	H, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	W, _ := strconv.Atoi(scanner.Text())

	A := make([][]int, H)
	rowSum := make([]int, H)
	colSum := make([]int, W)
	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
		for j := 0; j < W; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.Atoi(scanner.Text())
			rowSum[i] += A[i][j]
			colSum[j] += A[i][j]
		}
	}
	output := bufio.NewWriter(os.Stdout)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// Calculate B[i][j] as the sum of rowSum[i] and colSum[j] minus A[i][j] (because it's counted twice)
			if j > 0 {
				fmt.Fprint(output, " ")
			}
			fmt.Fprint(output, rowSum[i]+colSum[j]-A[i][j])
		}
		fmt.Fprintln(output)
	}
	output.Flush()

}
