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
	var n, t int
	fmt.Fscan(in, &n, &t)
	var s string
	fmt.Fscan(in, &s)
	q := []byte(s)
	for ; t > 0; t-- {
		i := 0
		for i < n-1 {
			if q[i] == byte('B') && q[i+1] == byte('G') {
				q[i] = byte('G')
				q[i+1] = byte('B')
				i += 2
			} else {
				i += 1
			}
		}
	}
	fmt.Fprint(out, string(q))
}
