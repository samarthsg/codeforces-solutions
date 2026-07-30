package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m, k int
		fmt.Fscan(in, &n, &m, &k)
		Lists := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &Lists[i])
		}
		known := make([]bool, n+1)
		for i := 0; i < k; i++ {
			var x int
			fmt.Fscan(in, &x)
			known[x] = true
		}
		if n == k {
			s := strings.Repeat("1", m)
			fmt.Fprintln(out, s)
			continue
		} else if k < n-1 {
			s := strings.Repeat("0", m)
			fmt.Fprintln(out, s)
			continue
		} else {
			var unknown int
			for i := 1; i <= n; i++ {
				if !known[i] {
					unknown = i
					break
				}
			}
			for _, missing := range Lists {
				if unknown == missing {
					fmt.Fprint(out, "1")
				} else {
					fmt.Fprint(out, "0")
				}
			}
		}
		fmt.Fprintln(out)
	}
}
