package main

import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	for ; t > 0; t--{
		var k , n int
		var s string
		fmt.Scan(&n, &k, &s)

		if k == n{
			for i := 0; i < n; i++{
				fmt.Print("-")
			}
			fmt.Println()
			continue
		}
		var zero, one, two int
		for i := 0; i < len(s); i++{
			if s[i] == '0' {
				zero ++
			} else if s[i] == '1' {
				one ++
			} else {
				two ++
			}
		}
		for i:= 1; i <= n; i++{
			if i <= zero || i > n-one{
				fmt.Print("-")
			}else if i >= zero+two+1 && i<=n-one-two{
				fmt.Print("+")
			}else{
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}
