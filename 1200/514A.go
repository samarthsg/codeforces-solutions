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
	var n, num int64
	fmt.Fscan(in, &n)
	var multiplier int64 = 1
	for n >= 10 {
		digit := n % 10
		if digit >= 5 {
			digit = 9 - digit
		}
		num = num + (digit * multiplier)
		multiplier *= 10
		n /= 10
	}
	digit := n % 10
	if digit != 9 && digit >= 5 {
		digit = 9 - digit
	}
	num = num + (digit * multiplier)
	fmt.Fprintln(out, num)
}
