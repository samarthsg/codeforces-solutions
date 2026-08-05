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
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if a == b || c == b || c == a {
			fmt.Fprintln(out, 0)
			continue
		}
		ac, bc, cc := a, b, c
		cnt := 0
		for {
			if a == min(ac, bc, cc) {
				a++
			} else if b == min(ac, bc, cc) {
				b++
			} else {
				c++
			}
			if a == max(ac, bc, cc) {
				a--
			} else if b == max(ac, bc, cc) {
				b--
			} else {
				c--
			}
			ac, bc, cc = a, b, c
			cnt++
			if a == b || b == c || c == a {
				break
			}
		}
		fmt.Fprintln(out, cnt)
	}
}
