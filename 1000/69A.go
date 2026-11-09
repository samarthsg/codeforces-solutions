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
	var x, y, z int
	fmt.Fscan(in, &n)
	for ; n > 0; n-- {
		var xi, yi, zi int
		fmt.Fscan(in, &xi, &yi, &zi)
		x += xi
		y += yi
		z += zi
	}
	if x == 0 && y == 0 && z == 0 {
		fmt.Fprint(out, "YES")
		return
	}
	fmt.Fprint(out, "NO")

}
