package main

import (
	"bufio"
	"fmt"
	"os"
)

type grp struct {
	ch    int
	value int
}

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
		var tracker []grp
		for _, num := range a {
			if len(tracker) > 0 && tracker[len(tracker)-1].ch == num {
				tracker[len(tracker)-1].value++
			} else {
				newE := grp{ch: num, value: 1}
				tracker = append(tracker, newE)
			}
		}
		if len(tracker) <= 1 {
			if tracker[0].value >= 2 {
				fmt.Fprintln(out, 1)
			} else {
				fmt.Fprintln(out, n)
			}
			continue
		}
		ans := n
		second := -1
		first := -1
		firstIndex := -1
		secondIndex := -1
		for i := 0; i < len(tracker); i++ {
			if tracker[i].value >= 2 && tracker[i].value >= first {
				second = first
				secondIndex = firstIndex
				first = tracker[i].value
				firstIndex = i
			}
		}
		if first != -1 && second != -1 && tracker[firstIndex].ch == tracker[secondIndex].ch {
			ans -= first
			fmt.Fprintln(out, ans)
			continue
		}
		if first != -1 && second != -1 {
			first--
			second--
			ans = ans - ((first + second) - 2)
			for i := 0; i < len(tracker); i++ {
				if i == firstIndex || i == secondIndex {
					continue
				} else if tracker[i].value >= 2 {
					ans = ans - (tracker[i].value - 1)
				}
			}
			fmt.Fprintln(out, ans)
			continue
		}
		if first != -1 {
			first--
			ans = ans - (first - 1)
			fmt.Fprintln(out, ans)
			continue
		}
		fmt.Fprintln(out, ans)
	}
}
