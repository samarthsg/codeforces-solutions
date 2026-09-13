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
		freq := make(map[int]int)
		var odds int
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
			if arr[i]%2 == 1 {
				odds++
			} else {
				freq[arr[i]%4]++
			}
		}
		ans := 0
		for _, value := range freq {
			if ans < value {
				ans = value
			}
		}
		if ans < odds {
			ans = odds
		}
		fmt.Fprintln(out, ans)
	}
}
