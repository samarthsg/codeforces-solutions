package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	n, row, coloum := 5, 0, 0
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &matrix[i][j])
			if matrix[i][j] == 1 {
				row, coloum = i, j
			}
		}
	}
	left := row - 2
	if left < 0 {
		left = -left
	}
	right := coloum - 2
	if right < 0 {
		right = -right
	}
	fmt.Fprintln(out, left+right)
}
