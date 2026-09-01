package main

import (
	"bufio"
	"fmt"
	"os"
)

func findmid(a int, b int, c int) int {
	if (b <= a && a <= c) || (c <= a && a <= b) {
		return a
	}
	if (a <= b && b <= c) || (c <= b && b <= a) {
		return b
	}
	return c
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		small := min(a, b, c)
		big := max(a, b, c)
		mid := findmid(a, b, c)
		for mid+small < big {
			big = mid + small
		}
		fmt.Fprintln(out, big-small)
	}
}
