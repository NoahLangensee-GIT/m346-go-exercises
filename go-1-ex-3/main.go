package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	//go run go-1-ex-3/main.go > eyes.txt
	//go run go-1-ex-3/main.go > dice.log
}
