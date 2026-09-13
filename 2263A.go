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
		ones, zeros := 0, 0
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			if a[i] == 1 {
				ones++
			} else {
				zeros++
			}
		}
		if zeros > ones {
			fmt.Fprintln(out, "Elsie")
			continue
		}
		fmt.Fprintln(out, "Bessie")
	}
}
