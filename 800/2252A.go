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
		freq, major := 0, -1
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
			if freq == 0 {
				major = arr[i]
				freq++
			} else if major == arr[i] {
				freq++
			} else {
				freq--
			}
		}
		freq = 0
		for i := 0; i < n; i++ {
			if arr[i] == major {
				freq++
			}
		}
		ans := 0
		gaps := freq - 2
		for i := 0; i < n; i++ {
			if major != arr[i] {
				gaps--
				ans += arr[i]
			}
		}
		if gaps > 0 {
			freq -= gaps
		}
		ans += freq * major
		fmt.Fprintln(out, ans)
	}
}
