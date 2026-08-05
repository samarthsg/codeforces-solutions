package main

import (
	"bufio"
	"fmt"
	"os"
)

type fs struct {
	ch   byte
	size int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		var s string
		fmt.Fscan(in, &n, &s)
		group := []fs{}
		cnt := 1
		a := s[0]
		for i := 1; i < n; i++ {
			if a == s[i] {
				cnt++
			} else {
				group = append(group, fs{
					ch:   a,
					size: cnt,
				})
				a = s[i]
				cnt = 1
			}
		}
		group = append(group, fs{
			ch:   a,
			size: cnt,
		})
		found := false
		ans := len(group)
		for i := 1; i < len(group)-1; i++ {
			if group[i].size == 1 && group[i-1].ch == group[i+1].ch {
				ans = len(group) - 2
				found = true
				break
			}
		}
		if found {
			fmt.Fprintln(out, ans)
			continue
		}
		for i := 1; i < len(group)-1; i++ {
			if group[i].size == 1 {
				ans = len(group) - 1
				break
			}
		}
		fmt.Fprintln(out, ans)
	}
}
