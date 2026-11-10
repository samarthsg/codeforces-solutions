package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var s string
	fmt.Fscan(in, &s)
	w := []byte(s)
	hash := make(map[string]int)
	hash["."] = 1
	hash["--"] = 3
	hash["-."] = 2
	res := ""
	var curr []byte
	for i := 0; i < len(w); i++ {
		curr = append(curr, w[i])
		if hash[string(curr)] != 0 {
			val := hash[string(curr)]
			val -= 1
			res += strconv.Itoa(val)
			curr = curr[:0]
		}
	}
	fmt.Fprintln(out, res)
}
