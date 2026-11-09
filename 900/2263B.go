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
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
		if k >= 2*n || k < n {
			fmt.Fprintln(out, -1)
			continue
		}
		m := 2*n - k
		put := m + 1
		dig := 1
		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j && dig <= m {
					matrix[i][j] = dig
					dig++
				} else {
					matrix[i][j] = put
					put++
				}
				fmt.Fprint(out, matrix[i][j], " ")
			}
			fmt.Fprintln(out)
		}
	}
}
