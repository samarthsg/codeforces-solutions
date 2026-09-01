package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, res int
	var statement string
	fmt.Fscan(in, &n)
	for ; n > 0; n-- {
		fmt.Fscan(in, &statement)
		if strings.Contains(statement, "++") {
			res++
		} else {
			res--
		}
	}
	fmt.Fprint(out, res)
}
