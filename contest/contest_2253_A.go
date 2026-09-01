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
		if n <= 2 {
			fmt.Fprintln(out, "YES")
			continue
		}
		if (n+1)%2 == 0 {
			fmt.Fprintln(out, "NO")
			continue
		}
		k := n + 1
		if k%2 == 0 || k%3 == 0 {
			fmt.Fprintln(out, "NO")
			continue
		}
		c := true
		for i := 5; i <= k/i; i += 6 {
			if k%i == 0 || k%(i+2) == 0 {
				c = false
				break
			}
		}
		if c == false {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
	}
}
