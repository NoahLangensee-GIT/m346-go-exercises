package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i <= Upper; i++ {
		var out string
		if i%3 == 0 && i%5 == 0 {
			out = "FizzBuzz"
		}
		fmt.Print(i, ". ")
	}
}
