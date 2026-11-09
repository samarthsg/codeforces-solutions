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
	var n, q int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	fmt.Fscan(in, &q)
	que := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &que[i])
	}
	hash := make(map[int]int, n)
	for i := 0; i < n; i++ {
		hash[arr[i]] = i + 1
	}
	var v, p int
	for i := 0; i < len(que); i++ {
		pos := hash[que[i]]
		v += pos
		p += (n - pos) + 1
	}
	fmt.Fprint(out, v, p)
}
