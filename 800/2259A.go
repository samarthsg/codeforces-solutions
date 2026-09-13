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
		cnt, cost := 0, 0
		done := false
		var s string
		fmt.Fscan(in, &s)
		for i := 0; i < n; i++ {
			cnt++
			if s[i] == '0' {
				done = true
			}
			if cnt >= k {
				cnt = 0
				if !done {
					cost++
				} else {
					done = false
				}
			}
		}
		fmt.Fprintln(out, cost)
	}
	fmt.Fprintln(out)
}
