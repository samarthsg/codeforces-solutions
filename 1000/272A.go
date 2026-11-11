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
	var n int
	fmt.Fscan(in, &n)
	f := make([]int, n)
	var t, ans int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &f[i])
		t += f[i]
	}
	for i := 1; i <= 5; i++ {
		if (t+i)%(n+1) != 1 {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
