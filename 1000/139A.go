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
	arr := make([]int, 7)
	for i := range 7 {
		fmt.Fscan(in, &arr[i])
	}
	i := 0
	for {
		t -= arr[i]
		if t <= 0 {
			break
		}
		i += 1
		i = i % 7
	}
	fmt.Fprintln(out, i+1)
}
