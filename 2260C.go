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
		var x, y int
		step, i := 0, 0
		fmt.Fscan(in, &x, &y)
		for {
			if x^y == x+y {
				step = i
				break
			}
			x--
			y++
			i++
		}
		fmt.Fprintln(out, x+y, step)
	}
}
