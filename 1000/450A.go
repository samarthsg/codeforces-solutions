package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	a   int
	pos int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	arr := make([]node, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i].a)
		arr[i].pos = i + 1
	}
	var ans int
	for len(arr) > 0 {
		n := arr[0]
		arr = arr[1:]
		n.a = n.a - m
		if n.a > 0 {
			arr = append(arr, n)
		} else {
			ans = n.pos
		}
	}
	fmt.Fprint(out, ans)
}
