package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	Go()
	Glass()
}
func Go() {
	fmt.Println(quote.Go())
}
func Glass() {
	fmt.Println(quote.Glass())
}
