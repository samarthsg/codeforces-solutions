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
		var x, y, k int
		fmt.Fscan(in, &x, &y, &k)
		d := y - x
		if d <= 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		ans := 0
		for k > 0 && d >= x {
			ans += d % x
			x++
			k--
		}
		if k > 0 {
			ans += k * d
		}
		fmt.Fprintln(out, ans)
	}
}
