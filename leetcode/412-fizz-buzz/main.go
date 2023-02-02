package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(20))
}

func fizzBuzz(n int) []string {
	out := make([]string, n)
	for index := range out {
		val := index + 1
		x := ""
		if (val%3) == 0 && (val%5) == 0 {
			x = "FizzBuzz"
		} else if (val % 3) == 0 {
			x = "Fizz"
		} else if (val % 5) == 0 {
			x = "Buzz"
		} else {
			x = strconv.Itoa(val)
		}
		out[index] = x
	}
	return out
}
