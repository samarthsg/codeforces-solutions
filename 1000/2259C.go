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
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		left, right := 0, n-1
		for left < n {
			if a[left] == 1 {
				break
			} else if a[left] == -1 {
				a[left] = 1
				break
			}
			left++
		}
		for right >= 0 {
			if a[right] == 1 {
				break
			}
			if a[right] == -1 {
				a[right] = 1
				break
			}
			right--
		}
		for i := 0; i < n; i++ {
			if a[i] == -1 {
				a[i] = 0
			}
			fmt.Fprint(out, a[i], " ")
		}
		fmt.Fprintln(out)
	}
}
