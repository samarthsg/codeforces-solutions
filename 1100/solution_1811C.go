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
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		b := make([]int, n-1)
		for i := 0; i < n-1; i++ {
			fmt.Fscan(in, &b[i])
		}
		a[0] = b[0]
		for i := 1; i < n-1; i++ {
			a[i] = min(b[i], b[i-1])
		}
		a[n-1] = b[n-2]

		for _, v := range a {
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)
	}
}
