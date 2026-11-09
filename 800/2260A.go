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
		zeros, ones := 0, 0
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			if a[i] == 0 {
				zeros++
			} else {
				ones++
			}
		}
		if a[0] == 0 && a[n-1] == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		if zeros < 2 {
			fmt.Fprintln(out, -1)
			continue
		}
		ans := 0
		if a[0] == 1 {
			ans++
		}
		if a[n-1] == 1 {
			ans++
		}
		fmt.Fprintln(out, ans)
	}
}
