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
		var x, y int
		fmt.Fscan(in, &x, &y)
		s := x + y
		for i := 1 << 30; i >= 1; i >>= 1 {
			if (s&i) != 0 && x >= i {
				x -= i
			}
		}
		fmt.Fprintln(out, s, x)
	}
}
