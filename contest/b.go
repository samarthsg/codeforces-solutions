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
		var s string
		fmt.Fscan(in, &n, &s)
		arr := []byte(s)
		ans := 0
		for i := 0; i < n; i++ {
			if ans == 4 || ans == 1 {
				break
			}
			if arr[i] != '?' {
				if i+2 < len(arr) && arr[i] == arr[i+2] {
					ans = -1
					break
				}
			} else {
				if i+2 < len(arr) {
					if arr[i] == arr[i+2] {
						arr[i] = '0'
						ans += 2
					} else if arr[i] != arr[i+2] && ans != 2 {
						ans += 1
					}
				} else {
					if i-2 >= 0 && arr[i-2] != arr[i] {
						ans += 2
					}
				}
			}
		}
		if ans == -1 {
			fmt.Fprintln(out, 0)
			continue
		}
		if ans == 0 {
			fmt.Fprintln(out, 1)
			continue
		}
		fmt.Println(out, ans)
	}
}
