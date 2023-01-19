package main

import (
	"fmt"

	"github.com/kr/pretty"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i < n; i++ {
		if i%3 != 0 && i%5 != 0 {
			pretty.Println(i)
			continue
		}
		if i%3 == 0 {
			pretty.Print("Fizz")
		}
		if i%5 == 0 {
			pretty.Print("Buzz")
		}
		pretty.Println("")
	}
}
