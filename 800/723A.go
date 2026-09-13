package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	nums := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &nums[i])
	}
	sort.Ints(nums)
	ans := nums[1] - nums[0] + nums[2] - nums[1]
	fmt.Fprintln(out, ans)
}
